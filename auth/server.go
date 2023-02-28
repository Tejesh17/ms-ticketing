package main

import (
	"example.com/auth/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.Routes(r)
	r.Run(":8080")
}
