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
  - name: counterName
    description: Name of the counter.
    type: string
    exposed: true
    required: true
    example_value: counter

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

  - name: value
    description: Value of the counter.
    type: integer
    exposed: true
    required: true
    example_value: 1
