package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	}
}
