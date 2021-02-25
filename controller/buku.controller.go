package controller

import "github.com/gin-gonic/gin"

//BukuController is ...
type BukuController interface {
	Create(c *gin.Context)
	FindByJudul(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
