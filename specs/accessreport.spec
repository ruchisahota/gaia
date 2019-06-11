# Model
model:
  rest_name: accessreport
  resource_name: accessreports
  entity_name: AccessReport
  package: zack
  group: policy/access
  description: Post a new access report.

# Attributes
attributes:
  v1:
  - name: action
    description: Action applied to the access.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - Accept
    - Reject
    example_value: Accept

  - name: claimHash
    description: Hash of the claims used to communicate.
    type: string
    exposed: true

  - name: content
    description: content of the report.
    type: string
    exposed: true
    required: true
    example_value: user X has tried to logged in

  - name: destinationID
    description: ID of the destination.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx

  - name: destinationIP
    description: Type of the destination.
    type: string
    exposed: true

  - name: destinationNamespace
    description: Namespace of the receiver.
    type: string
    exposed: true
    example_value: /my/namespace

  - name: destinationPort
    description: Port of the destination.
    type: integer
    exposed: true

  - name: destinationType
    description: Type of the source.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - ProcessingUnit
    - ExternalNetwork
    - Claims
    example_value: ProcessingUnit

  - name: dropReason
    description: |-
      This field is only set if 'action' is set to 'Reject' and specifies the reason
      for the rejection.
    type: string
    exposed: true

  - name: namespace
    description: This is here for backward compatibility.
    type: string
    exposed: true
    required: true
    deprecated: true
    example_value: /my/namespace

  - name: policyID
    description: ID of the policy that accepted the access.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx

  - name: policyNamespace
    description: Namespace of the policy that accepted the access.
    type: string
    exposed: true
    required: true
    example_value: /my/namespace

  - name: sourceID
    description: ID of the source.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx

  - name: sourceIP
    description: Type of the source.
    type: string
    exposed: true

  - name: sourceNamespace
    description: Namespace of the receiver.
    type: string
    exposed: true
    example_value: /my/namespace

  - name: sourceType
    description: Type of the source.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - ProcessingUnit
    - ExternalNetwork
    - Claims
    example_value: ProcessingUnit

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true

  - name: value
    description: Number of access in the report.
    type: integer
    exposed: true
    required: true
    example_value: 1
