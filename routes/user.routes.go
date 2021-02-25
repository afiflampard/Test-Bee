package routes

import (
	"helloworld/controller"
	"helloworld/middleware"

	"github.com/gin-gonic/gin"
)

//Router is ...
type Router interface {
	Routes(router *gin.Engine)
}

//RouterStruct is ...
type RouterStruct struct {
	userService        controller.UserController
	bukuService        controller.BukuController
	activityController controller.ActivityController
}

//NewUserController is ...
func NewUserController(userController controller.UserController, bukuController controller.BukuController, activityController controller.ActivityController) RouterStruct {
	return RouterStruct{userController, bukuController, activityController}
}

//Routes is ...
func (c *RouterStruct) Routes(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		v1.POST("/login", c.userService.Login)
		v1.POST("/signup", c.userService.SignUp)

		v1.Use(middleware.Authorization())
		v1.GET("/user/:id", c.userService.FindById)
		v1.GET("/user", c.userService.FindAll)
		v1.PUT("/user/:id", c.userService.Update)
		v1.DELETE("/user/:id", c.userService.Delete)
		v1.PUT("/user/:id/photo", c.userService.UpdatePhoto)

	}
	v2 := router.Group("/v1")
	{
		v3 := v2.Group("/buku")
		{
			v3.GET("/", c.bukuService.FindByJudul)
			v3.GET("/allBuku", c.bukuService.FindAll)
			v3.Use(middleware.Authorization())
			v3.POST("/:id", c.bukuService.Create)
			v3.PUT("/:id", c.bukuService.Update)
			v3.DELETE("/:id", c.bukuService.Delete)
		}
	}
	v4 := router.Group("/v1")
	{
		v5 := v4.Group("/activity")
		{
			v5.Use(middleware.Authorization())
			v5.POST("/pinjam/:id", c.activityController.PinjamBuku)
			v5.GET("/historypinjam/:id", c.activityController.HistoryPinjam)
			v5.GET("/historykembali/:id", c.activityController.HistoryKembali)
			v5.POST("/kembali/:id", c.activityController.KembaliBuku)

		}
	}

}
