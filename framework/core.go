package framework

import (
	"net/http"
	"strings"
)

type Core struct {
	router map[string]map[string]ControllerHandler
}

func NewCore() *Core {
	getRouter := map[string]ControllerHandler{}
	postRouter := map[string]ControllerHandler{}
	deleteRouter := map[string]ControllerHandler{}
	putRouter := map[string]ControllerHandler{}

	c := &Core{router: map[string]map[string]ControllerHandler{}}
	c.router["GET"] = getRouter
	c.router["POST"] = postRouter
	c.router["DELETE"] = deleteRouter
	c.router["PUT"] = putRouter

	return c
}

func (c *Core) Get(url string, handler ControllerHandler) {
	url = strings.ToUpper(url)
	c.router["GET"][url] = handler
}

func (c *Core) Post(url string, handler ControllerHandler) {
	url = strings.ToUpper(url)
	c.router["POST"][url] = handler
}
func (c *Core) Put(url string, handler ControllerHandler) {
	url = strings.ToUpper(url)
	c.router["PUT"][url] = handler
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	url = strings.ToUpper(url)
	c.router["DELETE"][url] = handler
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(w, r)

	router := c.FindRouteByRequest(r)
	if router == nil {
		ctx.JSON(404, "not found")
		return
	}

	if err := router(ctx); err != nil {
		ctx.JSON(500, "inner error")
		return
	}
}

func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	uri := request.URL.Path
	method := request.Method

	upperMethod := strings.ToUpper(method)
	upperURI := strings.ToUpper(uri)

	if methodHandlers, ok := c.router[upperMethod]; ok {
		if handler, ok := methodHandlers[upperURI]; ok {
			return handler
		}
	}

	return nil
}
