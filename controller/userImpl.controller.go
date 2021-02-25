package controller

import (
	"fmt"
	"helloworld/models"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

type loginController struct {
	jwtService JWTServices
}

//LoginHandler function
func LoginHandler(jwtServices JWTServices) UserController {
	return &loginController{jwtServices}
}

//SuccessResponse struct
type SuccessResponse struct {
	IsUser bool   `json:"isUser"`
	Token  string `json:"token"`
}

//ErrorResponse struct
type ErrorResponse struct {
	Error interface{} `json:"code"`
}

//InitiateDB db
func InitiateDB(dbData *gorm.DB) {
	db = dbData
}
func GetDB() *gorm.DB {
	return db
}

func (controller *loginController) Login(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "harus json ya")
	}
	password := u.Password
	if err := GetDB().Where("username = ?", u.Username).First(&u).Error; err != nil {
		Error(c, 401, "User not Found")
	}

	result := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if result != nil {
		Error(c, http.StatusUnauthorized, "Password not match")
	} else {
		token, err := controller.jwtService.GenerateToken(c, int(u.ID), u.Username, true)
		if err != nil {
			Error(c, http.StatusUnprocessableEntity, "Cannot Login")
		}
		c.JSON(200, &SuccessResponse{
			IsUser: true,
			Token:  token,
		})
	}

}

//Error handling function
func Error(c *gin.Context, code int, message interface{}) {
	payload := &ErrorResponse{
		Error: message,
	}
	c.JSON(code, payload)
}

func (controller *loginController) SignUp(c *gin.Context) {
	var person models.User

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "harus dimasukkan")
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(person.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	addPerson := models.User{
		Username: person.Username,
		Password: string(pass),
		FullName: person.FullName,
		Role:     person.Role,
	}
	err = GetDB().Debug().Create(&addPerson).Error
	if err != nil {
		c.JSON(401, &ErrorResponse{
			Error: err,
		})
	}

}

func (controller *loginController) FindById(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := GetDB().First(&user, id).Error; err != nil {
		Error(c, 404, "User Not Found")
	} else {
		c.JSON(200, &user)
	}

}

func (controller *loginController) FindAll(c *gin.Context) {
	var user []models.User

	if err := GetDB().Find(&user).Error; err != nil {
		Error(c, 404, "User Not Found")
	} else {
		c.JSON(200, &user)
	}
}

func (controller *loginController) Update(c *gin.Context) {

	var user models.User
	id := c.Param("id")
	if err := GetDB().First(&user, id).Error; err != nil {
		Error(c, 404, "User Not Found")
	}
	if user.Username != "" {
		var newUser models.User
		c.ShouldBindJSON(&newUser)
		pass, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.MinCost)
		user.Username = newUser.Username
		user.Password = string(pass)
		GetDB().Save(&user)
		c.JSON(200, "Data berhasil Di Update")

	} else {
		c.JSON(404, "User Not Found")
	}

}
func (controller *loginController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := GetDB().Delete(&models.User{}, id).Error; err != nil {
		Error(c, 404, "Not Delete")
	}
	c.JSON(200, "Coba")
}

func (controller *loginController) UpdatePhoto(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := GetDB().First(&user, id).Error; err != nil {
		Error(c, 404, "User Not Found")
	}
	file, header, err := c.Request.FormFile("file")
	filename := header.Filename
	out, err := os.Create("./tmp/" + filename)
	fmt.Print("Outnya", *out)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	user.Photo = out.Name()
	GetDB().Save(&user)
}
