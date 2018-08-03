# Model
model:
  rest_name: auth
  resource_name: auth
  entity_name: Auth
  package: midgard
  description: This API verifies if the given token is valid or not.

# Attributes
attributes:
  v1:
  - name: claims
    description: Claims are the claims.
    type: external
    exposed: true
    subtype: claims
    read_only: true
    autogenerated: true