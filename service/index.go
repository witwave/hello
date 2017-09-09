package service

import "github.com/gin-gonic/gin"

func (e *Env) Index(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "ping",
	})
}

