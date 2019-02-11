# Model
model:
  rest_name: token
  resource_name: tokens
  entity_name: Token
  package: barret
  group: internal/x509
  description: This api issue signed token from the given certificate.
  private: true

# Attributes
attributes:
  v1:
  - name: certificate
    description: Certificate contains the client certificate to use to create a token.
    type: string
    exposed: true
    required: true
    creation_only: true
    example_value: |-
      -----BEGIN CERTIFICATE-----
      MIIBsjCCAVigAwIBAgIQCjRniVLvvHhekhI3H7+m8DAKBggqhkjOPQQDAjBGMRAw
      DgYDVQQKEwdBcG9yZXRvMQ8wDQYDVQQLEwZhcG9tdXgxITAfBgNVBAMTGEFwb211
      eCBQdWJsaWMgU2lnbmluZyBDQTAeFw0xODAxMTcwNjUzNTFaFw0yNzExMjYwNjUz
      NTFaMDkxETAPBgNVBAoTCGFwb3Rlc3RzMREwDwYDVQQLEwhhcG90ZXN0czERMA8G
      A1UEAxMIYXBvdGVzdHMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASJsYjvdkaf
      qE/TL/eQETmIl8uFgqJi2LumYgEjHPHLlz8fJzDP12yQsqgOPymqmFCjmnVsdxcW
      YbXTwQzWW5kwozUwMzAOBgNVHQ8BAf8EBAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUH
      AwIwDAYDVR0TAQH/BAIwADAKBggqhkjOPQQDAgNIADBFAiEAgNxYC2hMqll3aTS8
      oBbFevZHAV5p0nyzRNj2pFrHEaQCIG4A9zKzE2f8g3zEaKcQNZ3Bxrk/T+hcJxaG
      Mlri6/+4
      -----END CERTIFICATE-----

  - name: signingKeyID
    description: SigningKeyID holds the ID of the custom CA to use to sign the token.
    type: string
    exposed: true
    stored: true

  - name: tags
    description: Tags includes a list of tags that must be added to the token.
    type: list
    exposed: true
    subtype: string
    creation_only: true

  - name: token
    description: Token contains the generated token.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: validity
    description: Validity contains the token validity duration.
    type: string
    exposed: true
    creation_only: true
