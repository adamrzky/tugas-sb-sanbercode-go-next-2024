package routes

import (
	"Quiz-1/controllers"
	"Quiz-1/middleware"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router) {
	router.GET("/bangun-datar/persegi", controllers.HandlePersegi)
	router.GET("/bangun-datar/persegi-panjang", controllers.HandlePersegiPanjang)
	router.GET("/bangun-datar/lingkaran", controllers.HandleLingkaran)
	router.GET("/bangun-ruang/kubus", controllers.HandleKubus)
	router.GET("/bangun-ruang/balok", controllers.HandleBalok)
	router.GET("/bangun-ruang/tabung", controllers.HandleTabung)

	router.GET("/categories", controllers.GetAllCategories)
	router.POST("/categories", middleware.BasicAuth(controllers.CreateCategory))
	router.PUT("/categories/:id", middleware.BasicAuth(controllers.UpdateCategory))
	router.DELETE("/categories/:id", middleware.BasicAuth(controllers.DeleteCategory))

	router.GET("/articles", controllers.GetAllArticles)
	router.POST("/articles", middleware.BasicAuth(controllers.CreateArticle))
	router.PUT("/articles/:id", middleware.BasicAuth(controllers.UpdateArticle))
	router.DELETE("/articles/:id", middleware.BasicAuth(controllers.DeleteArticle))

	router.GET("/categories/:id/articles", controllers.GetArticlesByCategory)
}
