# Model
model:
  rest_name: awsaccount
  resource_name: awsaccounts
  entity_name: AWSAccount
  package: vince
  group: core/authentication
  description: |-
    Allows you to bind an AWS account to your Aporeto account to allow
    auto-registration
    of enforcers running on EC2 instances.
  aliases:
  - aws
  - awsaccs
  - awsacc
  get:
    description: Retrieves the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@identifiable-stored'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: accessKeyID
    description: |-
      Contains the AWS access key ID. Aporeto uses this just to retrieve your
      account ID and does not store the value.
    type: string
    exposed: true
    required: true
    creation_only: true
    example_value: AKIAIOSFODNN7EXAMPLE

  - name: accessToken
    description: |-
      Contains your AWS access token. Aporeto uses this just to retrieve your
      account ID and does not store the value.
    type: string
    exposed: true
    creation_only: true

  - name: accountID
    description: Contains your verified account ID.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    filterable: true
    orderable: true

  - name: parentID
    description: Contains the parent Vince account ID.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    filterable: true

  - name: parentName
    description: Contains the name of the Vince parent account.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    filterable: true

  - name: region
    description: Contains your the region where your AWS account is located.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: us-west-2
    filterable: true
    orderable: true

  - name: secretAccessKey
    description: |-
      Contains the AWS secret access key. Aporeto uses this just to retrieve your
      account ID and does not store the value.
    type: string
    exposed: true
    required: true
    creation_only: true
    example_value: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
