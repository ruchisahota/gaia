{
    "attributes": [],
    "children": [],
    "model": {
        "aliases": [
            "hook"
        ],
        "create": null,
        "delete": true,
        "description": "Hook allows to to define hooks to the write operations in squall. Hooks are sent to an external Rufus server that will do the processing and eventually return a modified version of the object before we save it.",
        "entity_name": "HookPolicy",
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
        "package": "Policies",
        "resource_name": "hookpolicies",
        "rest_name": "hookpolicy",
        "root": null,
        "update": true
    }
}