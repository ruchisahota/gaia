{
    "attributes": [],
    "children": [],
    "model": {
        "aliases": [
            "quota"
        ],
        "create": null,
        "delete": true,
        "description": "Quotas Policies allows to set quotas on the number of objects that can be created in a namespace.",
        "entity_name": "QuotaPolicy",
        "extends": [
            "@base",
            "@described",
            "@disabled",
            "@identifiable-nopk-nostored",
            "@metadatable",
            "@named",
            "@propagated"
        ],
        "get": true,
        "package": "Policy",
        "resource_name": "quotapolicies",
        "rest_name": "quotapolicy",
        "root": null,
        "update": true
    }
}