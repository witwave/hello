package service

import (
	"github.com/gin-gonic/gin"
)

type Env struct {
	db string
}

func NewServer() *gin.Engine {
	r := gin.Default()

	env := &Env{db: "db"}
	r.GET("/", env.Index)
	return r
}
