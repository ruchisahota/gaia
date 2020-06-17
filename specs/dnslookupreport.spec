# Model
model:
  rest_name: dnslookupreport
  resource_name: dnslookupreports
  entity_name: DNSLookupReport
  package: zack
  group: policy/dns
  description: |-
    A DNS lookup report is used to report a DNS lookup that is happening on
    behalf of a processing unit. If the DNS server is on the standard UDP port 53
    then the defender can proxy the DNS traffic and make a report. The report
    indicate whether or not the lookup was successful.
  extends:
  - '@identifiable-stored'
  - '@zoned'
  - '@migratable'

# Attributes
attributes:
  v1:
  - name: action
    description: Action of the DNS request.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Accept
    - Reject
    example_value: Accept

  - name: enforcerID
    description: ID of the defender.
    type: string
    exposed: true
    stored: true

  - name: enforcerNamespace
    description: Namespace of the defender.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/namespace

  - name: processingUnitID
    description: ID of the PU.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx

  - name: processingUnitNamespace
    description: Namespace of the PU.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/namespace

  - name: reason
    description: |-
      This field is only set when the lookup fails. It specifies the reason for the
      failure.
    type: string
    exposed: true
    stored: true

  - name: resolvedName
    description: name used for DNS resolution.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: www.google.com

  - name: sourceIP
    description: Type of the source.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 10.0.0.1

  - name: timestamp
    description: Time and date of the log.
    type: time
    exposed: true
    stored: true

  - name: value
    description: Number of times the client saw this activity.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 1
