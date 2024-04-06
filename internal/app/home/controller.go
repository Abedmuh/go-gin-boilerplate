package home

import "github.com/gin-gonic/gin"

type CtrlInter interface {
	GetHome(c *gin.Context)
}

type CtrlImpl struct {}

func NewHomeController() *CtrlImpl {
	return &CtrlImpl{}
}

func (c *CtrlImpl) GetHome(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
    "message": "Hello World",
  })
}