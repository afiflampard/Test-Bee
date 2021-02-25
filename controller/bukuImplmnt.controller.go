package controller

import (
	"fmt"
	"helloworld/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type bukuImpl struct {
	jwtServices JWTServices
}

//Buku is ...
type Buku struct {
	KodeBuku    string `json:"kode_buku"`
	JudulBuku   string `json:"judul_buku"`
	PenulisBuku string `json:"penulis_buku"`
	Stok        uint   `json:"stok"`
}

//SuccessAddbuku is ...
type SuccessAddbuku struct {
	Kode    int    `json:"Status"`
	Message string `json:"message"`
}

//NewBukuController is ...
func NewBukuController(jwtservices JWTServices) BukuController {
	return &bukuImpl{jwtservices}
}

func (controller *bukuImpl) Create(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	// var role models.Role
	var u Buku
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "harus Json ya")
	}

	if err := GetDB().Model(&user).Preload("Role").Find(&user).First(&user, id).Error; err != nil {
		Error(c, 404, "User Not Found")
	}

	if strings.ToLower(user.Role.Role) == "petugas" {
		addBuku := models.Buku{
			KodeBuku:    u.KodeBuku,
			JudulBuku:   u.JudulBuku,
			PenulisBuku: u.PenulisBuku,
			Stok:        u.Stok,
			MaxStock:    u.Stok,
		}
		err := GetDB().Debug().Create(&addBuku).Error
		if err != nil {
			c.JSON(401, &ErrorResponse{
				Error: err,
			})
		} else {
			c.JSON(200, &SuccessAddbuku{
				Kode:    200,
				Message: "Success Add Buku",
			})
		}
	} else {
		Error(c, 404, "You Cannot add Buku")
	}

}

func (controller *bukuImpl) FindByJudul(c *gin.Context) {
	judulBuku := c.Query("judulBuku")
	var buku []models.Buku

	if err := GetDB().Find(&buku).Error; err != nil {
		Error(c, 404, "Buku not Found")
	}
	fmt.Println(buku)
	for _, bukuSearch := range buku {
		if strings.ToLower(bukuSearch.JudulBuku) == judulBuku {
			c.JSON(200, &bukuSearch)
		} else {
			Error(c, 404, "Buku not Found")
		}
	}

}

func (controller *bukuImpl) FindAll(c *gin.Context) {
	var buku []models.Buku

	if err := GetDB().Find(&buku).Error; err != nil {
		Error(c, 404, "Buku Not Found")
	} else {
		c.JSON(200, &buku)
	}
}

func (controller *bukuImpl) Update(c *gin.Context) {
	idBuku := c.Query("idBuku")
	fmt.Println(idBuku)
	idUser := c.Param("id")
	var user models.User
	var buku models.Buku
	var newBuku models.Buku

	if err := c.ShouldBindJSON(&newBuku); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Harus dimasukkan")
	}
	fmt.Println("Halooooß")

	if err := GetDB().Model(&user).Preload("Role").Find(&user).First(&user, idUser).Error; err != nil {
		Error(c, 404, "User Not Found")
	}

	if err := GetDB().First(&buku, idBuku).Error; err != nil {
		Error(c, 404, "Buku Not Found")
	}

	if strings.ToLower(user.Role.Role) == "petugas" {
		if newBuku.JudulBuku != "" {
			buku.JudulBuku = newBuku.JudulBuku
		} else if newBuku.PenulisBuku != "" {
			buku.PenulisBuku = newBuku.PenulisBuku
		} else if newBuku.JudulBuku != "" && newBuku.PenulisBuku != "" {
			buku.JudulBuku = newBuku.JudulBuku
			buku.PenulisBuku = newBuku.PenulisBuku
		} else {
			Error(c, 400, "Input Json")
		}
		err := GetDB().Save(&buku)
		if err != nil {
			c.JSON(200, "Data Berhasil Di Update")
		} else {
			c.JSON(400, "Data Tidak Berhasil Di Update")
		}
	}

}
func (controller *bukuImpl) Delete(c *gin.Context) {
	idUser := c.Param("id")
	idBuku := c.Query("idBuku")

	var user models.User

	if err := GetDB().Model(&user).Preload("Role").Find(&user).First(&user, idUser).Error; err != nil {
		Error(c, 404, "User Not Found")
	}

	if strings.ToLower(user.Role.Role) == "petugas" {
		if err := GetDB().Delete(&models.Buku{}, idBuku).Error; err != nil {
			Error(c, 404, "Not Delete")
		}
		c.JSON(200, "Deleted")
	} else {
		Error(c, 404, "You Cannot Delete Buku")
	}
}
