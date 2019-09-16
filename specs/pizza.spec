# Model
model:
  rest_name: pizza
  resource_name: pizzas
  entity_name: Pizza
  package: fry
  group: core/thing
  description: These are pizza.
  get:
    description: Retrieves one pizza.
  update:
    description: Updates one pizza.
  delete:
    description: Deletes one pizza.
  extends:
  - '@zoned'
  - '@base'
  - '@migratable'
  - '@identifiable-stored'
  - '@namespaced'

# Attributes
attributes:
  v1:
  - name: calories
    description: How fat you will get.
    type: integer
    exposed: true
    stored: true
    default_value: 400
    min_value: 100

  - name: name
    description: The name of pizza.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: Queen
