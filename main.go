package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bladewaltz9/Gee"
)

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	gee := Gee.New()
	gee.Use(Gee.Logger(), Gee.Recovery())

	gee.GET("/", func(c *Gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	gee.GET("/panic", func(c *Gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	gee.Run(":10000")
}
