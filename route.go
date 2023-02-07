package main

import "rainbow/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
