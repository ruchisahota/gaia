# Model
model:
  rest_name: privilegeescalationpolicy
  resource_name: privilegeescalationpolicies
  entity_name: PrivilegeEscalationPolicy
  package: squall
  group: policy/sudo
  description: |-
    The privilege escalation policy controls which PU/users can escalate privilege
    on a linux system.
  aliases:
  - privescpol
  - privescpols
  get:
    description: Retrives the PrivilegeEscalationPolicy with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the PrivilegeEscalationPolicy with the given ID.
  delete:
    description: Deletes the PrivilegeEscalationPolicy with the given ID.
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
  - '@fallback'
  - '@schedulable'
  - '@zonable'
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
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
