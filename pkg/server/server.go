package main

import (
	"fmt"
	"infra-practice/pkg/presentation/controller"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer(
	sampleController *controller.SampleController,
) *Server {
	router := &Server{
		Engine: gin.New(),
	}

	sample := router.Engine.Group("/sample")
	{
		sample.GET("/hello", sampleController.Hello)
	}

	return router
}

func (s *Server) Run(port string) error {
	if s == nil {
		panic("routerがnilです")
	}

	return s.Engine.Run(fmt.Sprintf(":%s", port))
}
