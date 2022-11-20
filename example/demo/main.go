package main

import (
	"lweb/pkg/jin"
	"net/http"
)

func main() {

	engine := jin.New()

	engine.Get("/", func(c *jin.Context) {
		c.String(http.StatusOK, "hello jin")
	})

	engine.Post("/user", func(c *jin.Context) {
		c.String(http.StatusOK, "user login")
	})

	engine.Get("/ljf/:id", func(c *jin.Context) {
		//ljfu
		c.JSON(http.StatusOK, jin.H{"name": "jin", "id": c.Param("id")})
	})

	engine.Run(":8099")

}
