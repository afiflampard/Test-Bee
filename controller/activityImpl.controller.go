package controller

import (
	"helloworld/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ActivityImpl struct {
	jwtServices JWTServices
}

type RequestPinjam struct {
	JudulBuku      string `json:"judul_buku"`
	TanggalKembali string `json:"tanggal_kembali"`
}

type SuccessPinjam struct {
	Kode    uint   `json:"status"`
	Message string `json:"message"`
}

func NewActivityController(jwtservices JWTServices) ActivityController {
	return &ActivityImpl{jwtservices}
}

const layoutFormat = "2006-01-02"

func (controller *ActivityImpl) PinjamBuku(c *gin.Context) {
	tx := GetDB().Begin()
	idMember := c.Param("id")
	idPetugas := c.Query("idPetugas")
	// fmt.Println(idMember)
	// fmt.Println(idPetugas)
	var buku []models.Buku
	var req RequestPinjam
	var petugas models.User
	var member models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Harus JSON ya")
	}

	var judulBuku []string
	if strings.Contains(req.JudulBuku, ",") {
		judulBuku = append(judulBuku, strings.Split(req.JudulBuku, ",")...)
	} else {
		judulBuku = append(judulBuku, req.JudulBuku)
	}

	if err := GetDB().Where("judul_buku IN ?", judulBuku).Find(&buku).Error; err != nil {
		Error(c, 404, "Buku Not Found")
	}
	if err := GetDB().Model(&petugas).Preload("Role").Find(&petugas, idPetugas).Error; err != nil {
		Error(c, 404, "Petugas Not Found")
	}

	if err := GetDB().Model(&member).Preload("Role").Find(&member, idMember).Error; err != nil {
		Error(c, 404, "Member Not Found")
	}

	t, _ := time.Parse(layoutFormat, req.TanggalKembali)

	if strings.ToLower(petugas.Role.Role) == "petugas" {
		pinjam := models.Order{
			TanggalPeminjaman: time.Now(),
			TanggalKembali:    t,
			IDPetugas:         petugas.ID,
			IDUser:            member.ID,
			NoState:           1,
		}
		err := GetDB().Debug().Create(&pinjam).Error
		if err != nil {
			c.JSON(401, &ErrorResponse{
				Error: err,
			})
			tx.Rollback()
		} else {
			tx.Commit()
			if len(buku) > 1 {
				for index := 0; index < len(buku); index++ {
					orderDetail := models.OrderDetail{
						IDOrder: pinjam.ID,
						IDBuku:  buku[index].ID,
					}
					err := GetDB().Debug().Create(&orderDetail).Error
					if err != nil {
						c.JSON(401, &ErrorResponse{
							Error: err,
						})
						tx.Rollback()
					}
					buku[index].Stok = buku[index].Stok - 1
					GetDB().Save(&buku)
					tx.Commit()
					history := models.History{
						IDBuku:  buku[index].ID,
						IDOrder: pinjam.ID,
						NoState: pinjam.NoState,
					}
					err = GetDB().Debug().Create(&history).Error
					if err != nil {
						c.JSON(401, &ErrorResponse{
							Error: err,
						})
						tx.Rollback()
					}
					tx.Commit()
				}

			} else {
				orderDetail := models.OrderDetail{
					IDOrder: pinjam.ID,
					IDBuku:  buku[0].ID,
				}
				err := GetDB().Debug().Create(&orderDetail).Error
				if err != nil {
					c.JSON(401, &ErrorResponse{
						Error: err,
					})
					tx.Rollback()
				}
				tx.Commit()
				buku[0].Stok = buku[0].Stok - 1
				GetDB().Save(&buku)
				history := models.History{
					IDBuku:  buku[0].ID,
					IDOrder: pinjam.ID,
					NoState: pinjam.NoState,
				}
				err = GetDB().Debug().Create(&history).Error
				if err != nil {
					c.JSON(401, &ErrorResponse{
						Error: err,
					})
					tx.Rollback()
				}
				tx.Commit()
			}
			c.JSON(200, &SuccessPinjam{
				Kode:    200,
				Message: "Buku Telah Dipinjam",
			})

		}
	}
}

func (controller *ActivityImpl) KembaliBuku(c *gin.Context) {
	tx := GetDB().Begin()

	idMember := c.Param("id")
	idPetugas := c.Query("idPetugas")
	var buku models.Buku
	var req RequestPinjam
	var petugas models.User
	var member models.User

	var orderDetail models.OrderDetail
	var order models.Order
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Harus JSON ya")
	}

	if err := GetDB().Where("judul_buku = ?", req.JudulBuku).First(&buku).Error; err != nil {
		Error(c, 404, "Buku Not Found")
	}
	if err := GetDB().Model(&petugas).Preload("Role").Find(&petugas, idPetugas).Error; err != nil {
		Error(c, 404, "Petugas Not Found")
	}

	if err := GetDB().Model(&member).Preload("Role").Find(&member, idMember).Error; err != nil {
		Error(c, 404, "Member Not Found")
	}
	if err := GetDB().Model(&orderDetail).Where("buku_id = ?", buku.ID).Preload("Order").Find(&orderDetail).Error; err != nil {
		Error(c, 404, "Buku Not Found")
	}
	if err := GetDB().Where("id = ?", orderDetail.IDOrder).First(&order).Error; err != nil {
		Error(c, 404, "Order Not Found")
	}
	if buku.Stok >= buku.MaxStock {
		Error(c, 400, "Buku Melebihi Max Stock")
	} else {
		if strings.ToLower(petugas.Role.Role) == "petugas" {
			orderDetail.Order.NoState = 2
			order.NoState = 2
			buku.Stok = buku.Stok + 1
			GetDB().Save(&order)
			GetDB().Save(&buku)
			history := models.History{
				IDBuku:  buku.ID,
				IDOrder: orderDetail.Order.ID,
				NoState: orderDetail.Order.NoState,
			}
			err := GetDB().Debug().Create(&history).Error
			if err != nil {
				c.JSON(401, &ErrorResponse{
					Error: err,
				})
				tx.Rollback()
			}
			tx.Commit()
			c.JSON(200, "Data Telah Terupdate")

		}
	}

}
func (controller *ActivityImpl) HistoryPinjam(c *gin.Context) {
	idPetugas := c.Param("id")

	var petugas models.User
	var histories []models.History
	if err := GetDB().Model(&petugas).Preload("Role").Find(&petugas, idPetugas).Error; err != nil {
		Error(c, 404, "Petugas Not Found")
	}
	if strings.ToLower(petugas.Role.Role) == "petugas" {
		if err := GetDB().Preload("Order.User.Role").Preload("Order.Petugas.Role").Preload("OrderState").Preload("Buku").Find(&histories).Error; err != nil {
			Error(c, 404, "History Not Found")
		} else {
			var tempHistory []models.History
			for _, history := range histories {
				if history.NoState == 1 {
					tempHistory = append(tempHistory, history)
					// c.JSON(200, history)
				}
			}
			c.JSON(200, tempHistory)
		}
	}
}

func (controller *ActivityImpl) HistoryKembali(c *gin.Context) {
	idPetugas := c.Param("id")

	var petugas models.User
	var histories []models.History
	if err := GetDB().Model(&petugas).Preload("Role").Find(&petugas, idPetugas).Error; err != nil {
		Error(c, 404, "Petugas Not Found")
	}
	if strings.ToLower(petugas.Role.Role) == "petugas" {
		if err := GetDB().Preload("Order.User.Role").Preload("Order.Petugas.Role").Preload("OrderState").Preload("Buku").Find(&histories).Error; err != nil {
			Error(c, 404, "History Not Found")
		} else {
			var tempHistory []models.History
			for _, history := range histories {
				if history.NoState == 2 {
					tempHistory = append(tempHistory, history)
				}
			}
			c.JSON(200, tempHistory)
		}
	}
}
