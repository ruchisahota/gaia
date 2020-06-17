# Model
model:
  rest_name: flowreport
  resource_name: flowreports
  entity_name: FlowReport
  package: zack
  group: policy/networking
  description: Post a new flow log.
  extends:
  - '@identifiable-stored'
  - '@zoned'
  - '@migratable'

# Indexes
indexes:
- - namespace
  - timestamp
- - timestamp
- - action
- - remoteNamespace
- - observedAction
- - dropReason
- - destinationPort
- - sourcePlatform
- - destinationPlatform

# Attributes
attributes:
  v1:
  - name: action
    description: Action applied to the flow.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Accept
    - Reject
    example_value: Accept

  - name: destinationController
    description: Identifier of the destination controller.
    type: string
    exposed: true
    example_value: api.east.acme.com

  - name: destinationID
    description: ID of the destination.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx

  - name: destinationIP
    description: Destination IP address.
    type: string
    exposed: true
    stored: true

  - name: destinationNamespace
    description: |-
      Namespace of the destination. This is deprecated. Use `remoteNamespace`. This
      property does nothing.
    type: string
    exposed: true
    stored: true
    deprecated: true
    example_value: /my/namespace
    omit_empty: true

  - name: destinationPlatform
    description: Identifier of the destination platform.
    type: string
    exposed: true
    stored: true
    example_value: api.east.acme.com

  - name: destinationPort
    description: Port of the destination.
    type: integer
    exposed: true
    stored: true

  - name: destinationType
    description: Destination type.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - ProcessingUnit
    - ExternalNetwork
    - Claims
    example_value: ProcessingUnit

  - name: dropReason
    description: |-
      This field is only set if `action` is set to `Reject`. It specifies the reason
      for the rejection.
    type: string
    exposed: true
    stored: true

  - name: encrypted
    description: If `true`, the flow was encrypted.
    type: boolean
    exposed: true
    stored: true

  - name: namespace
    description: This is here for backward compatibility.
    type: string
    exposed: true
    stored: true
    required: true
    deprecated: true
    example_value: /my/namespace

  - name: observed
    description: If `true`, design mode is on.
    type: boolean
    exposed: true
    stored: true

  - name: observedAction
    description: Action observed on the flow.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Accept
    - Reject
    - NotApplicable
    default_value: NotApplicable

  - name: observedDropReason
    description: |-
      Specifies the reason for a rejection. Only set if `observedAction` is set
      to `Reject`.
    type: string
    exposed: true
    stored: true

  - name: observedEncrypted
    description: Value of the encryption of the network policy that observed the flow.
    type: boolean
    exposed: true
    stored: true

  - name: observedPolicyID
    description: ID of the network policy that observed the flow.
    type: string
    exposed: true
    stored: true
    example_value: xxx-xxx-xxx

  - name: observedPolicyNamespace
    description: Namespace of the network policy that observed the flow.
    type: string
    exposed: true
    stored: true
    example_value: /my/namespace

  - name: policyID
    description: ID of the network policy that accepted the flow.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx

  - name: policyNamespace
    description: Namespace of the network policy that accepted the flow.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/namespace

  - name: protocol
    description: Protocol number.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 6

  - name: remoteNamespace
    description: Namespace of the object at the other end of the flow.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: serviceClaimHash
    description: Hash of the claims used to communicate.
    type: string
    exposed: true
    stored: true

  - name: serviceID
    description: ID of the service.
    type: string
    exposed: true
    stored: true

  - name: serviceNamespace
    description: Namespace of Service accessed.
    type: string
    exposed: true
    stored: true

  - name: serviceType
    description: ID of the service.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - L3
    - HTTP
    - TCP
    - NotApplicable
    default_value: NotApplicable

  - name: serviceURL
    description: Service URL accessed.
    type: string
    exposed: true
    stored: true

  - name: sourceController
    description: Identifier of the source controller.
    type: string
    exposed: true
    example_value: api.west.acme.com

  - name: sourceID
    description: ID of the source.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx

  - name: sourceIP
    description: Type of the source.
    type: string
    exposed: true
    stored: true

  - name: sourceNamespace
    description: |-
      Namespace of the source. This is deprecated. Use `remoteNamespace`. This
      property does nothing.
    type: string
    exposed: true
    stored: true
    deprecated: true
    example_value: /my/namespace
    omit_empty: true

  - name: sourcePlatform
    description: Identifier of the source platform.
    type: string
    exposed: true
    stored: true
    example_value: api.west.acme.com

  - name: sourceType
    description: Type of the source.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - ProcessingUnit
    - ExternalNetwork
    - Claims
    example_value: ProcessingUnit

  - name: timestamp
    description: Time and date of the log.
    type: time
    exposed: true
    stored: true

  - name: value
    description: Number of flows in the log.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 1
