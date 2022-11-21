package main

import (
	"fmt"
	"lweb/pkg/jin"
	"net/http"
)

func main() {

	engine := jin.Default()

	engine.Use(func(c *jin.Context) {
		fmt.Println("middleware 1")
		c.Next()
	})

	engine.Get("/", func(c *jin.Context) {
		fmt.Println("/  pre")
		c.Next()
	}, func(c *jin.Context) {
		fmt.Println("hello jin")
		c.String(http.StatusOK, "hello jin")
	})

	engine.Get("/ljf/:id", func(c *jin.Context) {
		//ljfu
		c.JSON(http.StatusOK, jin.H{"name": "jin", "id": c.Param("id")})
	})

	user := engine.Group("/user")
	{
		user.Post("/login", func(c *jin.Context) {
			c.String(http.StatusOK, "user login")
		})
	}

	engine.Run(":8099")

}
