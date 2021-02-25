package controller

import (
	"github.com/gin-gonic/gin"
)

//UserController interface
type UserController interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	UpdatePhoto(c *gin.Context)
}
