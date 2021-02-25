package controller

import "github.com/gin-gonic/gin"

type ActivityController interface {
	PinjamBuku(c *gin.Context)
	KembaliBuku(c *gin.Context)
	HistoryPinjam(c *gin.Context)
	HistoryKembali(c *gin.Context)
}
