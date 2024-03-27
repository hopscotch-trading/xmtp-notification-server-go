// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: notifications/v1/service.proto

// Protobuf Java Version: 4.26.0
package org.xmtp.android.library.push.notifications.v1;

public interface DeliveryMechanismOrBuilder extends
    // @@protoc_insertion_point(interface_extends:notifications.v1.DeliveryMechanism)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string apns_device_token = 1 [json_name = "apnsDeviceToken"];</code>
   * @return Whether the apnsDeviceToken field is set.
   */
  boolean hasApnsDeviceToken();
  /**
   * <code>string apns_device_token = 1 [json_name = "apnsDeviceToken"];</code>
   * @return The apnsDeviceToken.
   */
  java.lang.String getApnsDeviceToken();
  /**
   * <code>string apns_device_token = 1 [json_name = "apnsDeviceToken"];</code>
   * @return The bytes for apnsDeviceToken.
   */
  com.google.protobuf.ByteString
      getApnsDeviceTokenBytes();

  /**
   * <code>string firebase_device_token = 2 [json_name = "firebaseDeviceToken"];</code>
   * @return Whether the firebaseDeviceToken field is set.
   */
  boolean hasFirebaseDeviceToken();
  /**
   * <code>string firebase_device_token = 2 [json_name = "firebaseDeviceToken"];</code>
   * @return The firebaseDeviceToken.
   */
  java.lang.String getFirebaseDeviceToken();
  /**
   * <code>string firebase_device_token = 2 [json_name = "firebaseDeviceToken"];</code>
   * @return The bytes for firebaseDeviceToken.
   */
  com.google.protobuf.ByteString
      getFirebaseDeviceTokenBytes();

  /**
   * <code>string custom_token = 3 [json_name = "customToken"];</code>
   * @return Whether the customToken field is set.
   */
  boolean hasCustomToken();
  /**
   * <code>string custom_token = 3 [json_name = "customToken"];</code>
   * @return The customToken.
   */
  java.lang.String getCustomToken();
  /**
   * <code>string custom_token = 3 [json_name = "customToken"];</code>
   * @return The bytes for customToken.
   */
  com.google.protobuf.ByteString
      getCustomTokenBytes();

  org.xmtp.android.library.push.notifications.v1.DeliveryMechanism.DeliveryMechanismTypeCase getDeliveryMechanismTypeCase();
}
