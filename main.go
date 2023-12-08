package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/bladewaltz9/Gee"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	gee := Gee.New()
	gee.Use(Gee.Logger())

	gee.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	gee.LoadHTMLGlob("templates/*")
	gee.Static("/assets", "./static")

	stu1 := &student{Name: "tom", Age: 18}
	stu2 := &student{Name: "jack", Age: 20}

	gee.GET("/", func(c *Gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	gee.GET("/students", func(c *Gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", Gee.H{
			"title":  "Gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	gee.GET("/date", func(c *Gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", Gee.H{
			"title": "Gee",
			"now":   time.Date(2023, 12, 8, 0, 0, 0, 0, time.UTC),
		})
	})

	gee.Run(":10000")
}
