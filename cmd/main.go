package main

import (
	"chldu2000/paimon/internal/hello"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", hello.PaimonAlert)

	// Start the server on port 10086
	r.Run(":10086")
}
