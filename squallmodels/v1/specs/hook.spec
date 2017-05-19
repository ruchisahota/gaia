{
    "attributes": [],
    "children": [],
    "model": {
        "aliases": [],
        "create": null,
        "delete": true,
        "description": "Hook allows to to define hooks to the write operations in squall. Hooks are sent to an external Rufus server that will do the processing and eventually return a modified version of the object before we save it.",
        "entity_name": "Hook",
        "extends": [
            "@base",
            "@described"
        ],
        "get": true,
        "package": null,
        "resource_name": "hooks",
        "rest_name": "hook",
        "root": null,
        "update": true
    }
}