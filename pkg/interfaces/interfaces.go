package interfaces

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"time"

	v1 "github.com/xmtp/example-notification-server-go/pkg/proto/message_api/v1"
	"github.com/xmtp/example-notification-server-go/pkg/topics"
)

type DeliveryMechanismKind string

const (
	APNS DeliveryMechanismKind = "apns"
	FCM  DeliveryMechanismKind = "fcm"
)

type DeliveryMechanism struct {
	Kind      DeliveryMechanismKind
	Token     string
	UpdatedAt time.Time
}

type RegisterResponse struct {
	InstallationId string
	ValidUntil     time.Time
}

/*
*
An installation represents an app installed on a device. If the app is reinstalled, or installed onto
a new device it is expected to generate a fresh installation_id.
*/
type Installation struct {
	Id                string
	DeliveryMechanism DeliveryMechanism
}

type Subscription struct {
	Id             int64
	CreatedAt      time.Time
	InstallationId string
	Topic          string
	IsActive       bool
	IsSilent       bool
	HmacKey        *HmacKey
}

type SendRequest struct {
	IdempotencyKey string
	Message        *v1.Envelope
	MessageContext MessageContext
	Installation   Installation
	Subscription   Subscription
}

type MessageContext struct {
	MessageType topics.MessageType
	ShouldPush  *bool
	HmacInputs  *[]byte
	SenderHmac  *[]byte
}

func (m MessageContext) IsSender(hmacKey []byte) bool {
	if m.SenderHmac == nil || m.HmacInputs == nil {
		return false
	}
	hmacHash := hmac.New(sha256.New, hmacKey)
	hmacHash.Write(*m.HmacInputs)
	expectedHmac := hmacHash.Sum(nil)
	return hmac.Equal(*m.SenderHmac, expectedHmac)
}

type HmacKey struct {
	ThirtyDayPeriodsSinceEpoch int
	Key                        []byte
}

type SubscriptionInput struct {
	Topic    string
	IsSilent bool
	HmacKeys []HmacKey
}

type HmacUpdates map[string][]HmacKey

// Pluggable Installation Service interface
//
//go:generate mockery --dir ../interfaces --name Installations --output ../../mocks --outpkg mocks
type Installations interface {
	Register(ctx context.Context, installation Installation) (*RegisterResponse, error)
	Delete(ctx context.Context, installationId string) error
	GetInstallations(ctx context.Context, installationIds []string) ([]Installation, error)
}

// This interface is not expected to be pluggable
//
//go:generate mockery --dir ../interfaces --name Subscriptions --output ../../mocks --outpkg mocks
type Subscriptions interface {
	Subscribe(ctx context.Context, installationId string, topics []string) error
	Unsubscribe(ctx context.Context, installationId string, topics []string) error
	GetSubscriptions(ctx context.Context, topic string, thirtyDayPeriod int) ([]Subscription, error)
	SubscribeWithMetadata(ctx context.Context, installationId string, subscriptions []SubscriptionInput) error
}

// Pluggable interface for sending push notifications
type Delivery interface {
	CanDeliver(req SendRequest) bool
	Send(ctx context.Context, req SendRequest) error
}
