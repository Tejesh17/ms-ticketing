package routes

import (
	"example.com/auth/src/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(c *gin.Engine) {
	authroutes := c.Group("/auth")
	{
		authroutes.POST("/signup", handlers.SignupHander)
		authroutes.POST("/signin", handlers.SigninHandler)
	}

}
