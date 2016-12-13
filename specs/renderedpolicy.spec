{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "EgressPolicies lists all the egress policies attached to ProcessingUnit",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "egressPolicies",
            "orderable": true,
            "primary_key": false,
            "read_only": true,
            "required": false,
            "setter": false,
            "stored": false,
            "subtype": "rendered_policy",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "IngressPolicies lists all the ingress policies attached to ProcessingUnit",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "ingressPolicies",
            "orderable": true,
            "primary_key": false,
            "read_only": true,
            "required": false,
            "setter": false,
            "stored": false,
            "subtype": "rendered_policy",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Identifier of the ProcessingUnit",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": "free",
            "getter": false,
            "identifier": true,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "processingUnitID",
            "orderable": false,
            "primary_key": false,
            "read_only": true,
            "required": false,
            "setter": false,
            "stored": false,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": true,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Profile is the trust profile of the processing unit that should be used during all communications.",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "profile",
            "orderable": false,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "setter": null,
            "stored": true,
            "subtype": "trust_profile",
            "transient": false,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "create": false,
        "delete": false,
        "description": "RenderedPolicies attached to the given set of Subjects.",
        "entity_name": "RenderedPolicy",
        "extends": [
            "@identifiable-nopk-nostored"
        ],
        "get": false,
        "package": "Policies",
        "resource_name": "renderedpolicies",
        "rest_name": "renderedpolicy",
        "root": false,
        "update": false
    }
}