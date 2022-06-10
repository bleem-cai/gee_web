package main

import (
	"gee_web/gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello world\n")
	})
	r.Run(":8080")
}
