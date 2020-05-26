# Model
model:
  rest_name: healthcheck
  resource_name: healthchecks
  entity_name: HealthCheck
  package: ultros
  group: core/monitoring
  description: |-
    This API allows to retrieve a generic health state of the platform.
    A return code different from 200 OK means the platform is not operational.
    The health check contains the list of observed sub system.

# Attributes
attributes:
  v1:
  - name: alerts
    description: |-
      A human readable alert list describing the current state of the sub system if
      available.
    type: list
    exposed: true
    subtype: string
    read_only: true
    autogenerated: true
    omit_empty: true

  - name: name
    description: The name of the observed sub system if applicable.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    omit_empty: true

  - name: responseTime
    description: The response time of the observed sub system if applicable.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    omit_empty: true

  - name: status
    description: The current health of the observed sub system.
    type: enum
    exposed: true
    read_only: true
    allowed_choices:
    - Degraded
    - Offline
    - Operational
    autogenerated: true

  - name: type
    description: The type of the observed sub system.
    type: enum
    exposed: true
    read_only: true
    allowed_choices:
    - Cache
    - Database
    - General
    - MessagingSystem
    - Service
    - TSDB
    autogenerated: true