package jin

type IGroup interface {
	Get(string, HandleFunc)
	Post(string, HandleFunc)
	Put(string, HandleFunc)
	Delete(string, HandleFunc)

	Group(string) IGroup
}

type Group struct {
	engine *Engine
	prefix string
}

func NewGroup(e *Engine, uri string) IGroup {
	return &Group{
		engine: e,
		prefix: uri,
	}
}

func (g *Group) Get(uri string, handler HandleFunc) {
	g.engine.Get(g.prefix+uri, handler)
}

func (g *Group) Post(uri string, handler HandleFunc) {
	g.engine.Post(g.prefix+uri, handler)
}

func (g *Group) Put(uri string, handler HandleFunc) {
	g.engine.Put(g.prefix+uri, handler)
}

func (g *Group) Delete(uri string, handler HandleFunc) {
	g.engine.Delete(g.prefix+uri, handler)
}

func (g *Group) Group(prefix string) IGroup {
	return &Group{
		engine: g.engine,
		prefix: prefix,
	}
}
