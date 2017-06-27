---
title: Poke
keywords: api, Squall,
last_updated: June 9, 2017
tags: [Api,Squall]
summary: "When available, poke can be used to update various information about the parent. For instance, for enforcers, poke will be use as the heartbeat."
sidebar: squall_sidebar
permalink: squall_poke.html
folder: apidoc
---


## Attributes



## Code example

#### Retrieve all Pokes in a given Enforcer
<div class="ui top attached tabular menu">
  <a class="item active" data-tab="apoctl">apoctl</a>
  <a class="item" data-tab="golang">Go</a>
  <a class="item" data-tab="python">Python</a>
  <a class="item" data-tab="curl">curl</a>
</div>

```apoctl
apoctl api list poke in enforcer parent-id -n /namespace
```

```go
package main

import (
  "github.com/aporeto-inc/manipulate"
  "github.com/aporeto-inc/manipulate/manipwebsocket"

  squallmodels "github.com/aporeto-inc/gaia/squallmodels/current/golang"
)

func main() {

  m, disconnect, err := manipwebsocket.NewWebSocketManipulator("username", "token", "url", "namespace")

  if err != nil {
      panic(err)
  }

  var objects squallmodels.PokesList

  parent := squallmodels.NewEnforcer()
  parent.SetIdentifier("parent-id")

  mctx := manipulate.NewContext()
  mctx.Parent = parent

  if err = m.RetrieveMany(mctx, &amp;objects); err != nil {
    panic(err)
  }

  disconnect()
}
```

```python
import os

from squallmodels import Poke, Enforcer
from pymanipulate.context import Context
from pymanipulate.maniphttp import HTTPManipulator, PyManipulateError


def main():
    """
    """
    certificate = ("cert.pem", "key.pem")
    manipulator = HTTPManipulator("url", "/namespace", "username", "password", None, certificate)

    obj = Poke()
    objects = []

    parent = Enforcer(ID="parent-id")
    mctx = Context(parent=parent)

    try:
        manipulator.retrieve_children(mctx, obj.identity(), objects)
    except PyManipulateError as e:
        print "Err: %s" % e.errors[0]
        os._exit(1)

    for obj in objects:
        print obj.to_dict()

    exit(0)


if __name__ == "__main__":
    main()
```

```shell
curl -X GET \
  -H "X-Namespace: namespace" \
  -H "Authorization: token" \
  http://&lt;host&gt;:&lt;port&gt;/enforcers/id/poke
```



