package controller

import (
	"infra-practice/pkg/presentation/response"

	"github.com/gin-gonic/gin"
)

type SampleController struct{}

func (c *SampleController) Hello(ctx *gin.Context) {
	ctx.JSON(
		200,
		response.HelloResponse{
			Hello: "HelloWorld!!",
		},
	)
}

func NewSampleController() *SampleController {
	return &SampleController{}
}
