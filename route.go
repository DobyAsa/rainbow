package main

import "rainbow/framework"

func registerRouter(core *framework.Core) {
	core.Get("/user/login", UserLoginController)

	subjectApi := core.Group("subject")
	{
		subjectApi.Get("/get", UserLoginController)
	}
}
