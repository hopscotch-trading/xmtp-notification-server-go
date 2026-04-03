package api

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/xmtp/example-notification-server-go/mocks"
	"github.com/xmtp/example-notification-server-go/pkg/interfaces"
	"github.com/xmtp/example-notification-server-go/pkg/logging"
	"github.com/xmtp/example-notification-server-go/pkg/options"
	proto "github.com/xmtp/example-notification-server-go/pkg/proto/notifications/v1"
	protoconnect "github.com/xmtp/example-notification-server-go/pkg/proto/notifications/v1/notificationsv1connect"
)

const INSTALLATION_ID = "install1"

type testContext struct {
	cleanup           func()
	client            protoconnect.NotificationsClient
	ctx               context.Context
	httpClient        *http.Client
	installationsMock *mocks.Installations
	subscriptionsMock *mocks.Subscriptions
	apiServer         *ApiServer
}

func setupTest(t *testing.T) testContext {
	ctx := context.Background()
	port := getFreePort(t)
	installationsMock := mocks.NewInstallations(t)
	subscriptionsMock := mocks.NewSubscriptions(t)
	httpClient := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
	apiServer := NewApiServer(logging.CreateLogger("console", "info"), options.ApiOptions{Port: port}, installationsMock, subscriptionsMock)
	apiServer.Start()
	time.Sleep(50 * time.Millisecond)

	cleanup := func() {
		httpClient.CloseIdleConnections()
		apiServer.Stop()
	}

	return testContext{
		cleanup:           cleanup,
		client:            protoconnect.NewNotificationsClient(httpClient, fmt.Sprintf("http://127.0.0.1:%d", port)),
		ctx:               ctx,
		httpClient:        httpClient,
		installationsMock: installationsMock,
		subscriptionsMock: subscriptionsMock,
		apiServer:         apiServer,
	}
}

func getFreePort(t *testing.T) int {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	defer listener.Close()

	return listener.Addr().(*net.TCPAddr).Port
}

func Test_RegisterInstallation(t *testing.T) {
	ctx := setupTest(t)
	defer ctx.cleanup()

	deviceToken := "foo"
	validUntil := time.Now()

	ctx.installationsMock.On(
		"Register",
		mock.Anything,
		mock.Anything,
	).Return(&interfaces.RegisterResponse{
		InstallationId: INSTALLATION_ID,
		ValidUntil:     validUntil,
	}, nil)

	result, err := ctx.client.RegisterInstallation(
		ctx.ctx,
		connect.NewRequest(&proto.RegisterInstallationRequest{
			InstallationId: INSTALLATION_ID,
			DeliveryMechanism: &proto.DeliveryMechanism{
				DeliveryMechanismType: &proto.DeliveryMechanism_ApnsDeviceToken{ApnsDeviceToken: deviceToken},
			},
		}),
	)

	require.NoError(t, err)
	require.Equal(t, result.Msg.InstallationId, INSTALLATION_ID)
	require.Equal(t, result.Msg.ValidUntil, uint64(validUntil.UnixMilli()))
}

func Test_RegisterInstallationError(t *testing.T) {
	ctx := setupTest(t)
	defer ctx.cleanup()

	ctx.installationsMock.On(
		"Register",
		mock.Anything,
		mock.Anything,
	).Return(nil, errors.New("err"))

	result, err := ctx.client.RegisterInstallation(
		ctx.ctx,
		connect.NewRequest(&proto.RegisterInstallationRequest{
			InstallationId: INSTALLATION_ID,
			DeliveryMechanism: &proto.DeliveryMechanism{
				DeliveryMechanismType: &proto.DeliveryMechanism_ApnsDeviceToken{ApnsDeviceToken: "foo"},
			},
		}),
	)

	require.Equal(t, err.Error(), "internal: err")
	require.Nil(t, result)
}

func Test_DeleteInstallation(t *testing.T) {
	ctx := setupTest(t)
	defer ctx.cleanup()

	ctx.installationsMock.On("Delete", mock.Anything, mock.Anything).
		Return(nil)

	_, err := ctx.client.DeleteInstallation(
		ctx.ctx,
		connect.NewRequest(&proto.DeleteInstallationRequest{
			InstallationId: INSTALLATION_ID,
		}),
	)

	require.NoError(t, err)
	ctx.installationsMock.AssertCalled(
		t,
		"Delete",
		mock.Anything,
		INSTALLATION_ID,
	)
}

func Test_Subscribe(t *testing.T) {
	ctx := setupTest(t)
	defer ctx.cleanup()
	topics := []string{"topic1"}

	ctx.subscriptionsMock.On(
		"Subscribe",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)

	_, err := ctx.client.Subscribe(
		ctx.ctx,
		connect.NewRequest(&proto.SubscribeRequest{
			InstallationId: INSTALLATION_ID,
			Topics:         topics,
		}),
	)

	require.NoError(t, err)
	ctx.subscriptionsMock.AssertCalled(
		t,
		"Subscribe",
		mock.Anything,
		INSTALLATION_ID,
		topics,
	)
}

func Test_SubscribeError(t *testing.T) {
	ctx := setupTest(t)
	defer ctx.cleanup()

	ctx.subscriptionsMock.On(
		"Subscribe",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("test"))

	_, err := ctx.client.Subscribe(
		ctx.ctx,
		connect.NewRequest(&proto.SubscribeRequest{
			InstallationId: INSTALLATION_ID,
			Topics:         []string{"topic1"},
		}),
	)

	require.Error(t, err)
	require.Equal(t, err.Error(), "internal: test")
}

func Test_Unsubscribe(t *testing.T) {
	ctx := setupTest(t)
	defer ctx.cleanup()
	topics := []string{"topic1"}

	ctx.subscriptionsMock.On(
		"Unsubscribe",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)

	_, err := ctx.client.Unsubscribe(
		ctx.ctx,
		connect.NewRequest(&proto.UnsubscribeRequest{
			InstallationId: INSTALLATION_ID,
			Topics:         topics,
		}),
	)

	require.NoError(t, err)
	ctx.subscriptionsMock.AssertCalled(
		t,
		"Unsubscribe",
		mock.Anything,
		INSTALLATION_ID,
		topics,
	)
}
