package main

import (
	"net/http"

	"github.com/bladewaltz9/Gee"
)

func main() {
	gee := Gee.New()

	gee.GET("/", func(c *Gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := gee.Group("/v1")
	{
		v1.GET("/", func(c *Gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *Gee.Context) {
			// expect /hello?name=tom
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := gee.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *Gee.Context) {
			// expect /hello/tom
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		v2.POST("/login", func(c *Gee.Context) {
			c.JSON(http.StatusOK, Gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	gee.Run(":8080")
}
