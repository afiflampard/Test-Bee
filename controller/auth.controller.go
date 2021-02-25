package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//JWTServices interface
type JWTServices interface {
	GenerateToken(c *gin.Context, ID int, Username string, isUser bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
