package framework

type Group interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)
}

type DefaultGroup struct {
	core   *Core
	prefix string
}

func NewGroup(core *Core, prefix string) *DefaultGroup {
	return &DefaultGroup{
		core:   core,
		prefix: prefix,
	}
}

func (g *DefaultGroup) Get(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Get(uri, handler)
}

func (g *DefaultGroup) Post(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Post(uri, handler)
}

func (g *DefaultGroup) Put(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Put(uri, handler)
}

func (g *DefaultGroup) Delete(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Delete(uri, handler)
}

func (c *Core) Group(prefix string) Group {
	return NewGroup(c, prefix)
}
