package routes

import (
	"gin-gorm/configs/app_config"
	"gin-gorm/controllers/controller_anime"
	"gin-gorm/controllers/controller_book"
	"gin-gorm/controllers/controller_user"
	"gin-gorm/controllers/file_controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {
	route := app
	
	route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)

	// Route User
	route.GET("/user", controller_user.GetAllUser)
	route.GET("/user/:id", controller_user.GetUserById)
	route.GET("/user/paginate", controller_user.GetAllUserPaginate)
	route.POST("/user", controller_user.StoreUser)
	route.PATCH("/user/:id", controller_user.UpdateUserById)
	route.DELETE("/user/:id", controller_user.DeleteUserById)

	// Route Book
	route.GET("/book", controller_book.GetAllBook)

	// Route File
	route.POST("/file", file_controller.HandleUploadFile)

	// Route Anime
	route.GET("/anime", controller_anime.GetAllAnime)
	route.GET("/anime/:id", controller_anime.GetAnimeById)
	route.POST("/anime", controller_anime.CreateAnime)
	route.PATCH("/anime/:id", controller_anime.UpdateAnimeById)
	route.DELETE("/anime/:id", controller_anime.DeleteAnimeById)
	route.GET("/anime/paginate", controller_anime.GetAllAnimePaginate)
}