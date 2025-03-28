// DO NOT EDIT.
// swift-format-ignore-file
// swiftlint:disable all
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: identity/associations/signature.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

/// Signing methods for identity associations

import Foundation
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

/// RecoverableEcdsaSignature for EIP-191 and V2 signatures
public struct Xmtp_Identity_Associations_RecoverableEcdsaSignature: @unchecked Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// 65-bytes [ R || S || V ], with recovery id as the last byte
  public var bytes: Data = Data()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

/// EdDSA signature for 25519
public struct Xmtp_Identity_Associations_RecoverableEd25519Signature: @unchecked Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// 64 bytes [R(32 bytes) || S(32 bytes)]
  public var bytes: Data = Data()

  /// 32 bytes
  public var publicKey: Data = Data()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

/// Smart Contract Wallet signature
public struct Xmtp_Identity_Associations_SmartContractWalletSignature: @unchecked Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// CAIP-10 string
  /// https://github.com/ChainAgnostic/CAIPs/blob/main/CAIPs/caip-10.md
  public var accountID: String = String()

  /// Specify the block number to verify the signature against
  public var blockNumber: UInt64 = 0

  /// The actual signature bytes
  public var signature: Data = Data()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

/// Passkey signature
public struct Xmtp_Identity_Associations_RecoverablePasskeySignature: @unchecked Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var publicKey: Data = Data()

  public var signature: Data = Data()

  public var authenticatorData: Data = Data()

  public var clientDataJson: Data = Data()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

/// An existing address on xmtpv2 may have already signed a legacy identity key
/// of type SignedPublicKey via the 'Create Identity' signature.
/// For migration to xmtpv3, the legacy key is permitted to sign on behalf of the
/// address to create a matching xmtpv3 installation key.
/// This signature type can ONLY be used for CreateXid and AddAssociation
/// payloads, and can only be used once in xmtpv3.
public struct Xmtp_Identity_Associations_LegacyDelegatedSignature: Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var delegatedKey: Xmtp_MessageContents_SignedPublicKey {
    get {return _delegatedKey ?? Xmtp_MessageContents_SignedPublicKey()}
    set {_delegatedKey = newValue}
  }
  /// Returns true if `delegatedKey` has been explicitly set.
  public var hasDelegatedKey: Bool {return self._delegatedKey != nil}
  /// Clears the value of `delegatedKey`. Subsequent reads from it will return its default value.
  public mutating func clearDelegatedKey() {self._delegatedKey = nil}

  public var signature: Xmtp_Identity_Associations_RecoverableEcdsaSignature {
    get {return _signature ?? Xmtp_Identity_Associations_RecoverableEcdsaSignature()}
    set {_signature = newValue}
  }
  /// Returns true if `signature` has been explicitly set.
  public var hasSignature: Bool {return self._signature != nil}
  /// Clears the value of `signature`. Subsequent reads from it will return its default value.
  public mutating func clearSignature() {self._signature = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _delegatedKey: Xmtp_MessageContents_SignedPublicKey? = nil
  fileprivate var _signature: Xmtp_Identity_Associations_RecoverableEcdsaSignature? = nil
}

/// A wrapper for all possible signature types
public struct Xmtp_Identity_Associations_Signature: Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// Must have two properties:
  /// 1. An identifier (address or public key) for the signer must either be
  ///    recoverable, or specified as a field.
  /// 2. The signer certifies that the signing payload is correct. The payload
  ///    must be inferred from the context in which the signature is provided.
  public var signature: Xmtp_Identity_Associations_Signature.OneOf_Signature? = nil

  public var erc191: Xmtp_Identity_Associations_RecoverableEcdsaSignature {
    get {
      if case .erc191(let v)? = signature {return v}
      return Xmtp_Identity_Associations_RecoverableEcdsaSignature()
    }
    set {signature = .erc191(newValue)}
  }

  public var erc6492: Xmtp_Identity_Associations_SmartContractWalletSignature {
    get {
      if case .erc6492(let v)? = signature {return v}
      return Xmtp_Identity_Associations_SmartContractWalletSignature()
    }
    set {signature = .erc6492(newValue)}
  }

  public var installationKey: Xmtp_Identity_Associations_RecoverableEd25519Signature {
    get {
      if case .installationKey(let v)? = signature {return v}
      return Xmtp_Identity_Associations_RecoverableEd25519Signature()
    }
    set {signature = .installationKey(newValue)}
  }

  public var delegatedErc191: Xmtp_Identity_Associations_LegacyDelegatedSignature {
    get {
      if case .delegatedErc191(let v)? = signature {return v}
      return Xmtp_Identity_Associations_LegacyDelegatedSignature()
    }
    set {signature = .delegatedErc191(newValue)}
  }

  public var passkey: Xmtp_Identity_Associations_RecoverablePasskeySignature {
    get {
      if case .passkey(let v)? = signature {return v}
      return Xmtp_Identity_Associations_RecoverablePasskeySignature()
    }
    set {signature = .passkey(newValue)}
  }

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  /// Must have two properties:
  /// 1. An identifier (address or public key) for the signer must either be
  ///    recoverable, or specified as a field.
  /// 2. The signer certifies that the signing payload is correct. The payload
  ///    must be inferred from the context in which the signature is provided.
  public enum OneOf_Signature: Equatable, Sendable {
    case erc191(Xmtp_Identity_Associations_RecoverableEcdsaSignature)
    case erc6492(Xmtp_Identity_Associations_SmartContractWalletSignature)
    case installationKey(Xmtp_Identity_Associations_RecoverableEd25519Signature)
    case delegatedErc191(Xmtp_Identity_Associations_LegacyDelegatedSignature)
    case passkey(Xmtp_Identity_Associations_RecoverablePasskeySignature)

  }

  public init() {}
}

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "xmtp.identity.associations"

extension Xmtp_Identity_Associations_RecoverableEcdsaSignature: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".RecoverableEcdsaSignature"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "bytes"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularBytesField(value: &self.bytes) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.bytes.isEmpty {
      try visitor.visitSingularBytesField(value: self.bytes, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Xmtp_Identity_Associations_RecoverableEcdsaSignature, rhs: Xmtp_Identity_Associations_RecoverableEcdsaSignature) -> Bool {
    if lhs.bytes != rhs.bytes {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Xmtp_Identity_Associations_RecoverableEd25519Signature: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".RecoverableEd25519Signature"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "bytes"),
    2: .standard(proto: "public_key"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularBytesField(value: &self.bytes) }()
      case 2: try { try decoder.decodeSingularBytesField(value: &self.publicKey) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.bytes.isEmpty {
      try visitor.visitSingularBytesField(value: self.bytes, fieldNumber: 1)
    }
    if !self.publicKey.isEmpty {
      try visitor.visitSingularBytesField(value: self.publicKey, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Xmtp_Identity_Associations_RecoverableEd25519Signature, rhs: Xmtp_Identity_Associations_RecoverableEd25519Signature) -> Bool {
    if lhs.bytes != rhs.bytes {return false}
    if lhs.publicKey != rhs.publicKey {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Xmtp_Identity_Associations_SmartContractWalletSignature: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".SmartContractWalletSignature"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "account_id"),
    2: .standard(proto: "block_number"),
    3: .same(proto: "signature"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.accountID) }()
      case 2: try { try decoder.decodeSingularUInt64Field(value: &self.blockNumber) }()
      case 3: try { try decoder.decodeSingularBytesField(value: &self.signature) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.accountID.isEmpty {
      try visitor.visitSingularStringField(value: self.accountID, fieldNumber: 1)
    }
    if self.blockNumber != 0 {
      try visitor.visitSingularUInt64Field(value: self.blockNumber, fieldNumber: 2)
    }
    if !self.signature.isEmpty {
      try visitor.visitSingularBytesField(value: self.signature, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Xmtp_Identity_Associations_SmartContractWalletSignature, rhs: Xmtp_Identity_Associations_SmartContractWalletSignature) -> Bool {
    if lhs.accountID != rhs.accountID {return false}
    if lhs.blockNumber != rhs.blockNumber {return false}
    if lhs.signature != rhs.signature {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Xmtp_Identity_Associations_RecoverablePasskeySignature: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".RecoverablePasskeySignature"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "public_key"),
    2: .same(proto: "signature"),
    3: .standard(proto: "authenticator_data"),
    4: .standard(proto: "client_data_json"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularBytesField(value: &self.publicKey) }()
      case 2: try { try decoder.decodeSingularBytesField(value: &self.signature) }()
      case 3: try { try decoder.decodeSingularBytesField(value: &self.authenticatorData) }()
      case 4: try { try decoder.decodeSingularBytesField(value: &self.clientDataJson) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.publicKey.isEmpty {
      try visitor.visitSingularBytesField(value: self.publicKey, fieldNumber: 1)
    }
    if !self.signature.isEmpty {
      try visitor.visitSingularBytesField(value: self.signature, fieldNumber: 2)
    }
    if !self.authenticatorData.isEmpty {
      try visitor.visitSingularBytesField(value: self.authenticatorData, fieldNumber: 3)
    }
    if !self.clientDataJson.isEmpty {
      try visitor.visitSingularBytesField(value: self.clientDataJson, fieldNumber: 4)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Xmtp_Identity_Associations_RecoverablePasskeySignature, rhs: Xmtp_Identity_Associations_RecoverablePasskeySignature) -> Bool {
    if lhs.publicKey != rhs.publicKey {return false}
    if lhs.signature != rhs.signature {return false}
    if lhs.authenticatorData != rhs.authenticatorData {return false}
    if lhs.clientDataJson != rhs.clientDataJson {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Xmtp_Identity_Associations_LegacyDelegatedSignature: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".LegacyDelegatedSignature"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "delegated_key"),
    2: .same(proto: "signature"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._delegatedKey) }()
      case 2: try { try decoder.decodeSingularMessageField(value: &self._signature) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._delegatedKey {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try { if let v = self._signature {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Xmtp_Identity_Associations_LegacyDelegatedSignature, rhs: Xmtp_Identity_Associations_LegacyDelegatedSignature) -> Bool {
    if lhs._delegatedKey != rhs._delegatedKey {return false}
    if lhs._signature != rhs._signature {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Xmtp_Identity_Associations_Signature: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".Signature"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "erc_191"),
    2: .standard(proto: "erc_6492"),
    3: .standard(proto: "installation_key"),
    4: .standard(proto: "delegated_erc_191"),
    5: .same(proto: "passkey"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try {
        var v: Xmtp_Identity_Associations_RecoverableEcdsaSignature?
        var hadOneofValue = false
        if let current = self.signature {
          hadOneofValue = true
          if case .erc191(let m) = current {v = m}
        }
        try decoder.decodeSingularMessageField(value: &v)
        if let v = v {
          if hadOneofValue {try decoder.handleConflictingOneOf()}
          self.signature = .erc191(v)
        }
      }()
      case 2: try {
        var v: Xmtp_Identity_Associations_SmartContractWalletSignature?
        var hadOneofValue = false
        if let current = self.signature {
          hadOneofValue = true
          if case .erc6492(let m) = current {v = m}
        }
        try decoder.decodeSingularMessageField(value: &v)
        if let v = v {
          if hadOneofValue {try decoder.handleConflictingOneOf()}
          self.signature = .erc6492(v)
        }
      }()
      case 3: try {
        var v: Xmtp_Identity_Associations_RecoverableEd25519Signature?
        var hadOneofValue = false
        if let current = self.signature {
          hadOneofValue = true
          if case .installationKey(let m) = current {v = m}
        }
        try decoder.decodeSingularMessageField(value: &v)
        if let v = v {
          if hadOneofValue {try decoder.handleConflictingOneOf()}
          self.signature = .installationKey(v)
        }
      }()
      case 4: try {
        var v: Xmtp_Identity_Associations_LegacyDelegatedSignature?
        var hadOneofValue = false
        if let current = self.signature {
          hadOneofValue = true
          if case .delegatedErc191(let m) = current {v = m}
        }
        try decoder.decodeSingularMessageField(value: &v)
        if let v = v {
          if hadOneofValue {try decoder.handleConflictingOneOf()}
          self.signature = .delegatedErc191(v)
        }
      }()
      case 5: try {
        var v: Xmtp_Identity_Associations_RecoverablePasskeySignature?
        var hadOneofValue = false
        if let current = self.signature {
          hadOneofValue = true
          if case .passkey(let m) = current {v = m}
        }
        try decoder.decodeSingularMessageField(value: &v)
        if let v = v {
          if hadOneofValue {try decoder.handleConflictingOneOf()}
          self.signature = .passkey(v)
        }
      }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    switch self.signature {
    case .erc191?: try {
      guard case .erc191(let v)? = self.signature else { preconditionFailure() }
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    }()
    case .erc6492?: try {
      guard case .erc6492(let v)? = self.signature else { preconditionFailure() }
      try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
    }()
    case .installationKey?: try {
      guard case .installationKey(let v)? = self.signature else { preconditionFailure() }
      try visitor.visitSingularMessageField(value: v, fieldNumber: 3)
    }()
    case .delegatedErc191?: try {
      guard case .delegatedErc191(let v)? = self.signature else { preconditionFailure() }
      try visitor.visitSingularMessageField(value: v, fieldNumber: 4)
    }()
    case .passkey?: try {
      guard case .passkey(let v)? = self.signature else { preconditionFailure() }
      try visitor.visitSingularMessageField(value: v, fieldNumber: 5)
    }()
    case nil: break
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Xmtp_Identity_Associations_Signature, rhs: Xmtp_Identity_Associations_Signature) -> Bool {
    if lhs.signature != rhs.signature {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
