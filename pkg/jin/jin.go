package jin

import (
	"log"
	"net/http"
)

type Engine struct {
	router *Router
}

type H map[string]interface{}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) Get(url string, handle HandleFunc) {
	e.router.addRoute("get", url, handle)
}

func (e *Engine) Post(url string, handle HandleFunc) {
	e.router.addRoute("post", url, handle)
}

func (e *Engine) Put(url string, handle HandleFunc) {
	e.router.addRoute("put", url, handle)
}

func (e *Engine) Delete(url string, handle HandleFunc) {
	e.router.addRoute("delete", url, handle)
}

func (e *Engine) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, e))
}

func (e *Engine) Group(prefix string) IGroup {
	return NewGroup(e, prefix)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	context := &Context{writer: w, request: r}
	//交给路由处理
	e.router.handle(context)
}
