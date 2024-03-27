// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: notifications/v1/service.proto

// Protobuf Java Version: 4.26.0
package org.xmtp.android.library.push.notifications.v1;

public interface SubscribeWithMetadataRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:notifications.v1.SubscribeWithMetadataRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string installation_id = 1 [json_name = "installationId"];</code>
   * @return The installationId.
   */
  java.lang.String getInstallationId();
  /**
   * <code>string installation_id = 1 [json_name = "installationId"];</code>
   * @return The bytes for installationId.
   */
  com.google.protobuf.ByteString
      getInstallationIdBytes();

  /**
   * <code>repeated .notifications.v1.Subscription subscriptions = 2 [json_name = "subscriptions"];</code>
   */
  java.util.List<org.xmtp.android.library.push.notifications.v1.Subscription> 
      getSubscriptionsList();
  /**
   * <code>repeated .notifications.v1.Subscription subscriptions = 2 [json_name = "subscriptions"];</code>
   */
  org.xmtp.android.library.push.notifications.v1.Subscription getSubscriptions(int index);
  /**
   * <code>repeated .notifications.v1.Subscription subscriptions = 2 [json_name = "subscriptions"];</code>
   */
  int getSubscriptionsCount();
  /**
   * <code>repeated .notifications.v1.Subscription subscriptions = 2 [json_name = "subscriptions"];</code>
   */
  java.util.List<? extends org.xmtp.android.library.push.notifications.v1.SubscriptionOrBuilder> 
      getSubscriptionsOrBuilderList();
  /**
   * <code>repeated .notifications.v1.Subscription subscriptions = 2 [json_name = "subscriptions"];</code>
   */
  org.xmtp.android.library.push.notifications.v1.SubscriptionOrBuilder getSubscriptionsOrBuilder(
      int index);
}
