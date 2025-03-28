// DO NOT EDIT.
// swift-format-ignore-file
// swiftlint:disable all
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: device_sync/consent_backup.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

/// Definitions for backups

import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that you are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

/// Consent record type
public enum Xmtp_DeviceSync_ConsentBackup_ConsentTypeSave: SwiftProtobuf.Enum, Swift.CaseIterable {
  public typealias RawValue = Int
  case unspecified // = 0
  case conversationID // = 1
  case inboxID // = 2

  /// NOTE: This enum value was marked as deprecated in the .proto file
  case address // = 3
  case UNRECOGNIZED(Int)

  public init() {
    self = .unspecified
  }

  public init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .unspecified
    case 1: self = .conversationID
    case 2: self = .inboxID
    case 3: self = .address
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  public var rawValue: Int {
    switch self {
    case .unspecified: return 0
    case .conversationID: return 1
    case .inboxID: return 2
    case .address: return 3
    case .UNRECOGNIZED(let i): return i
    }
  }

  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static let allCases: [Xmtp_DeviceSync_ConsentBackup_ConsentTypeSave] = [
    .unspecified,
    .conversationID,
    .inboxID,
    .address,
  ]

}

/// Consent record state
public enum Xmtp_DeviceSync_ConsentBackup_ConsentStateSave: SwiftProtobuf.Enum, Swift.CaseIterable {
  public typealias RawValue = Int
  case unspecified // = 0
  case unknown // = 1
  case allowed // = 2
  case denied // = 3
  case UNRECOGNIZED(Int)

  public init() {
    self = .unspecified
  }

  public init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .unspecified
    case 1: self = .unknown
    case 2: self = .allowed
    case 3: self = .denied
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  public var rawValue: Int {
    switch self {
    case .unspecified: return 0
    case .unknown: return 1
    case .allowed: return 2
    case .denied: return 3
    case .UNRECOGNIZED(let i): return i
    }
  }

  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static let allCases: [Xmtp_DeviceSync_ConsentBackup_ConsentStateSave] = [
    .unspecified,
    .unknown,
    .allowed,
    .denied,
  ]

}

/// Proto representation of a consent record save
public struct Xmtp_DeviceSync_ConsentBackup_ConsentSave: Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var entityType: Xmtp_DeviceSync_ConsentBackup_ConsentTypeSave = .unspecified

  public var state: Xmtp_DeviceSync_ConsentBackup_ConsentStateSave = .unspecified

  public var entity: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "xmtp.device_sync.consent_backup"

extension Xmtp_DeviceSync_ConsentBackup_ConsentTypeSave: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "CONSENT_TYPE_SAVE_UNSPECIFIED"),
    1: .same(proto: "CONSENT_TYPE_SAVE_CONVERSATION_ID"),
    2: .same(proto: "CONSENT_TYPE_SAVE_INBOX_ID"),
    3: .same(proto: "CONSENT_TYPE_SAVE_ADDRESS"),
  ]
}

extension Xmtp_DeviceSync_ConsentBackup_ConsentStateSave: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "CONSENT_STATE_SAVE_UNSPECIFIED"),
    1: .same(proto: "CONSENT_STATE_SAVE_UNKNOWN"),
    2: .same(proto: "CONSENT_STATE_SAVE_ALLOWED"),
    3: .same(proto: "CONSENT_STATE_SAVE_DENIED"),
  ]
}

extension Xmtp_DeviceSync_ConsentBackup_ConsentSave: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ConsentSave"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "entity_type"),
    2: .same(proto: "state"),
    3: .same(proto: "entity"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularEnumField(value: &self.entityType) }()
      case 2: try { try decoder.decodeSingularEnumField(value: &self.state) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.entity) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.entityType != .unspecified {
      try visitor.visitSingularEnumField(value: self.entityType, fieldNumber: 1)
    }
    if self.state != .unspecified {
      try visitor.visitSingularEnumField(value: self.state, fieldNumber: 2)
    }
    if !self.entity.isEmpty {
      try visitor.visitSingularStringField(value: self.entity, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Xmtp_DeviceSync_ConsentBackup_ConsentSave, rhs: Xmtp_DeviceSync_ConsentBackup_ConsentSave) -> Bool {
    if lhs.entityType != rhs.entityType {return false}
    if lhs.state != rhs.state {return false}
    if lhs.entity != rhs.entity {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
