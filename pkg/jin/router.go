package jin

import (
	"net/http"
	"strings"
)

var methodType = [4]string{"get", "post", "put", "delete"}

type Router struct {
	handlers map[string]*Tree
}

func newRouter() *Router {
	router := &Router{handlers: make(map[string]*Tree)}

	for _, v := range methodType {
		router.handlers[strings.ToUpper(v)] = NewTree()
	}

	return router
}

func (r *Router) addRoute(method string, url string, handler HandleFunc) {
	if tree, ok := r.handlers[strings.ToUpper(method)]; ok {
		tree.Set(url, 0, handler)
	}
}

func (r *Router) getRoute(method string, url string) (map[string]string, HandleFunc) {
	if tree, ok := r.handlers[strings.ToUpper(method)]; ok {
		pattern, handler := tree.Get(url)
		//解析路径参数
		params := r.parseUriParam(url, pattern)
		return params, handler
	}
	return nil, nil
}

func (r *Router) parseUriParam(url string, pattern string) map[string]string {
	p := strings.Split(pattern, "/")
	segments := strings.Split(url, "/")

	params := make(map[string]string)
	for i := 0; i < len(p); i++ {
		if strings.HasPrefix(p[i], ":") {
			params[p[i][1:]] = segments[i]
		}
	}
	return params
}

//根据请求，执行对应的处理函数
func (r *Router) handle(c *Context) {
	method := c.request.Method
	if _, ok := r.handlers[strings.ToUpper(method)]; !ok {
		c.String(http.StatusForbidden, strings.Join([]string{"The [", method, "]", "is not supported"}, " "))
		return
	}

	url := c.request.URL.Path
	//寻找路由
	params, handleFunc := r.getRoute(method, url)
	c.params = params

	if handleFunc == nil {
		c.String(http.StatusNotFound, "404 not found")
		return
	}
	//执行对应处理函数
	handleFunc(c)
}
