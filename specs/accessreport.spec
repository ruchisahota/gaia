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

  - name: enforcerID
    description: Identifier of the enforcer.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx

  - name: enforcerNamespace
    description: Namespace of the enforcer.
    type: string
    exposed: true
    required: true
    example_value: /my/namespace

  - name: puID
    description: ID of the PU.
    type: string
    exposed: true
    example_value: xxx-xxx-xxx

  - name: puNamespace
    description: Namespace of the PU.
    type: string
    exposed: true
    example_value: /my/namespace

  - name: reason
    description: |-
      This field is only set if 'action' is set to 'Reject' and specifies the reason
      for the rejection.
    type: string
    exposed: true

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true

  - name: type
    description: Type of the report.
    type: string
    exposed: true
    required: true
    allowed_choices:
    - SSHLogIn
    - SSHLogOut
    - SudoLogIn
    - SudoLogOut
    default_value:
    - SSHLogIn

  - name: value
    description: Number of access in the report.
    type: integer
    exposed: true
    required: true
    example_value: 1
