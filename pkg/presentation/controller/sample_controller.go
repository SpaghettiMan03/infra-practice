package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type SampleController struct{}

func (c *SampleController) Hello(ctx *gin.Context) {
	fmt.Println("Hello!")
}

func NewSampleController() *SampleController {
	return &SampleController{}
}
