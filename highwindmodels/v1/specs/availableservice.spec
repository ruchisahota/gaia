{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "CategoryID of the service.",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "categoryID",
            "orderable": false,
            "primary_key": null,
            "read_only": true,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": null,
            "type": "string",
            "unique": null,
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
            "description": "Description of the service",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "description",
            "orderable": false,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": null,
            "type": "string",
            "unique": null,
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
            "description": "Icon contains a base64 image for the available service. ",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "icon",
            "orderable": false,
            "primary_key": null,
            "read_only": true,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": null,
            "type": "string",
            "unique": null,
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
            "description": "LongDescription contains a more detailed description of the service.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "longDescription",
            "orderable": false,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": null,
            "type": "string",
            "unique": null,
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
            "description": "Name of the Service",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "name",
            "orderable": false,
            "primary_key": false,
            "read_only": null,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": null,
            "type": "string",
            "unique": false,
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
            "description": "Parameters of the service the user can or has to specify",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "parameters",
            "orderable": false,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": false,
            "subtype": "service_parameters",
            "transient": null,
            "type": "external",
            "unique": null,
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
            "description": "Title represents the title of the service.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "title",
            "orderable": false,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": null,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "aliases": [
            "asrv"
        ],
        "create": null,
        "delete": false,
        "description": "AvailableService represents a service that is available for launching",
        "entity_name": "AvailableService",
        "extends": [],
        "get": false,
        "package": null,
        "resource_name": "availableservices",
        "rest_name": "availableservice",
        "root": null,
        "update": false
    }
}