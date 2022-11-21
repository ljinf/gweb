package jin

import (
	"fmt"
	"testing"
)

func add() *Router {

	router := newRouter()
	router.addRoute("get", "/", func(c *Context) {
		fmt.Println("get /")
	})
	router.addRoute("get", "/user", func(c *Context) {
		fmt.Println("get /user")
	})
	router.addRoute("get", "/ljf/login", func(c *Context) {
		fmt.Println("get /user/login")
	})

	router.addRoute("get", "/ljf/:id", func(c *Context) {
		fmt.Println("get /ljf/:id")
	})

	router.addRoute("post", "/user/:id", func(c *Context) {
		fmt.Println("post /user/:id")
	})
	router.addRoute("post", "/user/:id/name", func(c *Context) {
		fmt.Println("post /user/:id/name")
	})
	router.addRoute("post", "/user/:id/age", func(c *Context) {
		fmt.Println("post /user/:id/age")
	})

	return router
}

func TestRouter(t *testing.T) {
	router := add()

	_, handle := router.getRoute("get", "/user")
	//_, handle := router.getRoute("post", "/user/1")
	if handle != nil {
		handle(nil)
	} else {
		fmt.Println("404 not found")
	}
}
