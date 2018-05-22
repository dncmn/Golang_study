package router

import (
	"github.com/gin-gonic/gin"
	"self/handler"
)

func Router(r *gin.Engine) {

	user := r.Group("/user")
	user.GET("/name", handler.GetUserNameHandler)
	user.GET("/health_check", handler.ConsulHealthCheck)
	r.Run()

}

// /user/health_check
