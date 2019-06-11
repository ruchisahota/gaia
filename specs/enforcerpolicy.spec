# Model
model:
  rest_name: enforcerpolicy
  resource_name: enforcerpolicies
  entity_name: EnforcerPolicy
  package: squall
  group: policy/sudo
  description: The enforcer policy controls who can access to the host.
  aliases:
  - epol
  - epols
  get:
    description: Retrives the EnforcerPolicy with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the EnforcerPolicy with the given ID.
  delete:
    description: Deletes the EnforcerPolicy with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@described'
  - '@disabled'
  - '@identifiable-not-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@schedulable'
  - '@zonable'
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: allowSudoAccess
    description: AllowSudoAccess indicates if the user is allowed to use sudo commands.
    type: boolean
    exposed: true

  - name: object
    description: |-
      Object contains the tag expression matching the enforcers the subject is allowed
      to connect to.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

  - name: subject
    description: |-
      Subject contains the tag expression the tags need to match for the policy to
      apply.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression
