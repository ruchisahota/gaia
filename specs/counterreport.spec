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
  - name: CounterName
    description: Name of the counter.
    type: string
    exposed: true
    required: true
    example_value: counter

  - name: destinationIP
    description: DestinationIP is the IP address of the destination.
    type: string
    exposed: true

  - name: destinationPort
    description: DestinationPort is the destination port of a TCP or UDP counter.
    type: integer
    exposed: true
    example_value: 11000
    max_value: 65536

  - name: dropReason
    description: |-
      This field is only set if 'event' is set to 'Dropped' and specifies the reason
      for the drop.
    type: string
    exposed: true

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
    required: true
    example_value: xxx-xxx-xxx
    filterable: true

  - name: processingUnitNamespace
    description: Namespace of the PU reporting the counter.
    type: string
    exposed: true
    required: true
    example_value: /my/namespace
    filterable: true

  - name: timestamp
    description: Timestamp is the date of the report.
    type: time
    exposed: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"

  - name: type
    description: Type of counter.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - Received
    - Transmitted
    - Dropped
    example_value: Rcv
