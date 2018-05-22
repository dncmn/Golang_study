package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserNameHandler(c *gin.Context) {
	fmt.Println("hello")
	c.JSON(http.StatusOK, gin.H{"name": "tom"})
}

func ConsulHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
