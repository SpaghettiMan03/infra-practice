package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer() *Server {
	router := &Server{
		Engine: gin.New(),
	}

	sample := router.Engine.Group("/sample")
	{
		sample.GET("/hello", func(ctx *gin.Context) {

		})
	}

	return router
}

func (s *Server) Run(port string) error {
	if s == nil {
		panic("routerがnilです")
	}

	return s.Engine.Run(fmt.Sprintf(":%s", port))
}
