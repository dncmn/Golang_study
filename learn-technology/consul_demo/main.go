package main

import (
	"github.com/gin-gonic/gin"
	"self/router"
)

func main() {
	router.Router(gin.Default())
}
