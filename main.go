package main

import (
	"net/http"

	"github.com/bladewaltz9/Gee"
)

func main() {
	gee := Gee.New()

	gee.GET("/", func(c *Gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	gee.GET("/hello", func(c *Gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	gee.GET("/hello/:name", func(c *Gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	gee.GET("/hello/:name/:age", func(c *Gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s, your age is %s\n", c.Param("name"), c.Path, c.Param("age"))
	})

	gee.GET("/assets/*filepath", func(c *Gee.Context) {
		c.JSON(http.StatusOK, Gee.H{"filepath": c.Param("filepath")})
	})

	gee.POST("/login", func(c *Gee.Context) {
		c.JSON(http.StatusOK, Gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	gee.Run(":8080")
}
