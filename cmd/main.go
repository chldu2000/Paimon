package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define a route for the root URL
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "前面的区域，以后再来探索吧！")
	})

	// Start the server on port 10086
	r.Run(":10086")
}
