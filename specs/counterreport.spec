# Model
model:
  rest_name: counterreport
  resource_name: counterreports
  entity_name: CounterReport
  package: zack
  group: core/enforcer
  description: Post a new counter tracing report.

# Attributes
attributes:
  v1:
  - name: counterAckInUnknownState
    description: Counter for sending finack ack received in uknown connection state.
    type: integer
    exposed: true
    default_value: 0

  - name: counterAckInvalidFormat
    description: Counter for ack packet dropped because of invalid format.
    type: integer
    exposed: true
    default_value: 0

  - name: counterAckRejected
    description: Counter for reject ack packet as per policy.
    type: integer
    exposed: true
    default_value: 0

  - name: counterAckSigValidationFailed
    description: Counter for ack packet dropped because signature validation failed.
    type: integer
    exposed: true
    default_value: 0

  - name: counterAckTCPNoTCPAuthOption
    description: Counter for tcp authentication option not found.
    type: integer
    exposed: true
    default_value: 0

  - name: counterConnectionsProcessed
    description: Counter for ".
    type: integer
    exposed: true
    default_value: 0

  - name: counterContextIDNotFound
    description: Counter for unable to find contextid.
    type: integer
    exposed: true
    default_value: 0

  - name: counterDroppedExternalService
    description: |-
      Counter for no acls found for external services. dropping application syn
      packet.
    type: integer
    exposed: true
    default_value: 0

  - name: counterInvalidConnState
    description: Counter for invalid connection state.
    type: integer
    exposed: true
    default_value: 0

  - name: counterInvalidNetState
    description: Counter for invalid net state.
    type: integer
    exposed: true
    default_value: 0

  - name: counterInvalidProtocol
    description: Counter for invalid protocol.
    type: integer
    exposed: true
    default_value: 0

  - name: counterInvalidSynAck
    description: Counter for pu is already dead - drop synack packet.
    type: integer
    exposed: true
    default_value: 0

  - name: counterMarkNotFound
    description: Counter for pu mark not found.
    type: integer
    exposed: true
    default_value: 0

  - name: counterNetSynNotSeen
    description: Counter for network syn packet was not seen.
    type: integer
    exposed: true
    default_value: 0

  - name: counterNoConnFound
    description: Counter for no context or connection found.
    type: integer
    exposed: true
    default_value: 0

  - name: counterNonPUTraffic
    description: Counter for traffic that belongs to a non PU process.
    type: integer
    exposed: true
    default_value: 0

  - name: counterOutOfOrderSynAck
    description: Counter for synack for flow with processed finack.
    type: integer
    exposed: true
    default_value: 0

  - name: counterPortNotFound
    description: Counter for port not found.
    type: integer
    exposed: true
    default_value: 0

  - name: counterRejectPacket
    description: Counter for reject the packet as per policy.
    type: integer
    exposed: true
    default_value: 0

  - name: counterServicePostprocessorFailed
    description: Counter for post service processing failed for network packet.
    type: integer
    exposed: true
    default_value: 0

  - name: counterServicePreprocessorFailed
    description: Counter for pre service processing failed for network packet.
    type: integer
    exposed: true
    default_value: 0

  - name: counterSynAckBadClaims
    description: Counter for synack packet dropped because of bad claims.
    type: integer
    exposed: true
    default_value: 0

  - name: counterSynAckClaimsMisMatch
    description: Counter for syn/ack packet dropped because of encryption mismatch.
    type: integer
    exposed: true
    default_value: 0

  - name: counterSynAckDroppedExternalService
    description: Counter for synack from external service dropped.
    type: integer
    exposed: true
    default_value: 0

  - name: counterSynAckInvalidFormat
    description: Counter for synack packet dropped because of invalid format.
    type: integer
    exposed: true
    default_value: 0

  - name: counterSynAckMissingClaims
    description: Counter for synack packet dropped because of no claims.
    type: integer
    exposed: true
    default_value: 0

  - name: counterSynAckMissingToken
    description: Counter for synack packet dropped because of missing token.
    type: integer
    exposed: true
    default_value: 0

  - name: counterSynAckNoTCPAuthOption
    description: Counter for tcp authentication option not found.
    type: integer
    exposed: true
    default_value: 0

  - name: counterSynAckRejected
    description: Counter for dropping because of reject rule on transmitter.
    type: integer
    exposed: true
    default_value: 0

  - name: counterSynDroppedInvalidFormat
    description: Counter for syn packet dropped because of invalid format.
    type: integer
    exposed: true
    default_value: 0

  - name: counterSynDroppedInvalidToken
    description: Counter for syn packet dropped because of invalid token.
    type: integer
    exposed: true
    default_value: 0

  - name: counterSynDroppedNoClaims
    description: Counter for syn packet dropped because of no claims.
    type: integer
    exposed: true
    default_value: 0

  - name: counterSynDroppedTCPOption
    description: Counter for tcp authentication option not found.
    type: integer
    exposed: true
    default_value: 0

  - name: counterSynRejectPacket
    description: Counter for syn dropped due to policy.
    type: integer
    exposed: true
    default_value: 0

  - name: counterSynUnexpectedPacket
    description: Counter for received syn packet from unknown pu.
    type: integer
    exposed: true
    default_value: 0

  - name: counterTCPAuthNotFound
    description: Counter for tcp authentication option not found.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPAckInvalidSignature
    description: Counter for dropped udp ack invalid signature.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPConnectionsProcessed
    description: Counter for number of processed UDP connections.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPDropContextNotFound
    description: Counter for dropped UDP data packets with no context.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPDropFin
    description: Counter for dropped udp FIN handshake packets.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPDropInNfQueue
    description: Counter for dropped UDP in NfQueue.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPDropNoConnection
    description: Counter for dropped UDP data packets with no connection.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPDropPacket
    description: Counter for dropped UDP data packets.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPDropQueueFull
    description: Counter for dropped UDP Queue Full.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPDropSynAck
    description: Counter for dropped udp synack handshake packets.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPInvalidNetState
    description: Counter for udp packets received in invalid network state.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPPostProcessingFailed
    description: Counter for UDP packets failing postprocessing.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPPreProcessingFailed
    description: Counter for UDP packets failing preprocessing.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPRejected
    description: Counter for UDP packets dropped due to policy.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPSynAckDropBadClaims
    description: Counter for dropped udp synack bad claims.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPSynAckMissingClaims
    description: Counter for dropped udp synack missing claims.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPSynAckPolicy
    description: Counter for dropped udp synack bad claims.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPSynDrop
    description: Counter for dropped udp syn transmits.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPSynDropPolicy
    description: Counter for dropped udp syn policy.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPSynInvalidToken
    description: Counter for dropped udp FIN handshake packets.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUDPSynMissingClaims
    description: Counter for dropped UDP SYN missing claims.
    type: integer
    exposed: true
    default_value: 0

  - name: counterUnknownError
    description: Counter for unknown error.
    type: integer
    exposed: true
    default_value: 0

  - name: enforcerID
    description: Identifier of the enforcer sending the report.
    type: string
    exposed: true
    stored: true
    default_value: xxxx-xxx-xxxx

  - name: enforcerNamespace
    description: Namespace of the enforcer sending the report.
    type: string
    exposed: true
    stored: true
    default_value: /my/namespace

  - name: processingUnitID
    description: PUID is the ID of the PU reporting the counter.
    type: string
    exposed: true
    example_value: xxx-xxx-xxx
    filterable: true

  - name: processingUnitNamespace
    description: Namespace of the PU reporting the counter.
    type: string
    exposed: true
    example_value: /my/namespace
    filterable: true

  - name: timestamp
    description: Timestamp is the date of the report.
    type: time
    exposed: true
    example_value: "2018-06-14T23:10:46.420397985Z"
