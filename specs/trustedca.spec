# Model
model:
  rest_name: trustedca
  resource_name: trustedcas
  entity_name: TrustedCA
  package: squall
  group: policy/enforcerconfig
  description: Represents a trusted certificate authority (CA).
  aliases: []

# Attributes
attributes:
  v1:
  - name: certificate
    description: The namespace X.509 CA the enforcer should use.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: type
    description: Identifies the type of CA.
    type: enum
    exposed: true
    read_only: true
    allowed_choices:
    - X509
    - SSH
    autogenerated: true