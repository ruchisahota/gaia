# Model
model:
  rest_name: enforcerreport
  resource_name: enforcerreports
  entity_name: EnforcerReport
  package: zack
  group: core/enforcer
  description: Post a new enforcer statistics report.
  extends:
  - '@identifiable-stored'
  - '@zoned'
  - '@migratable'

# Attributes
attributes:
  v1:
  - name: CPULoad
    description: Total CPU utilization of the defender as a percentage of vCPUs.
    type: float
    exposed: true
    stored: true
    example_value: 10

  - name: ID
    description: ID of the defender.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: memory
    description: Total resident memory used by the defender in bytes.
    type: integer
    exposed: true
    stored: true
    example_value: 10000

  - name: name
    description: Name of the defender.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: aporeto-enforcerd-xxx

  - name: namespace
    description: Namespace of the defender.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/ns

  - name: processes
    description: Number of active processes of the defender.
    type: integer
    exposed: true
    stored: true
    example_value: 10

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true
    stored: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"
