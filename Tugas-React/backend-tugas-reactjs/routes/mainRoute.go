package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"backend-tugas-reactjs/controllers"
	"backend-tugas-reactjs/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB, r *gin.Engine) {

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	r.Use(cors.New(corsConfig))

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/movies", controllers.GetAllMovie)
	r.GET("/movies/:id", controllers.GetMovieById)

	r.GET("/age-rating-categories", controllers.GetAllRating)
	r.GET("/age-rating-categories/:id", controllers.GetRatingById)
	r.GET("/age-rating-categories/:id/movies", controllers.GetMoviesByRatingId)

	moviesMiddlewareRoute := r.Group("/movies")
	moviesMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	moviesMiddlewareRoute.POST("", controllers.CreateMovie)
	moviesMiddlewareRoute.PUT("/:id", controllers.UpdateMovie)
	moviesMiddlewareRoute.DELETE("/:id", controllers.DeleteMovie)

	ratingMiddlewareRoute := r.Group("/age-rating-categories")
	ratingMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	ratingMiddlewareRoute.POST("", controllers.CreateRating)
	ratingMiddlewareRoute.PUT("/:id", controllers.UpdateRating)
	ratingMiddlewareRoute.DELETE("/:id", controllers.DeleteRating)

	r.GET("/footballers", controllers.GetAllFootballer)
	r.POST("/footballers", controllers.CreateFootballer)
	r.GET("/footballers/:id", controllers.GetFootballerById)
	r.PUT("/footballers/:id", controllers.UpdateFootballer)
	r.DELETE("/footballers/:id", controllers.DeleteFootballer)

	// r.GET("/books", controllers.GetAllBooks)
	// r.POST("/books", controllers.CreateBook)
	// r.GET("/books/:id", controllers.GetBookByID)
	// r.PUT("/books/:id", controllers.UpdateBook)
	// r.DELETE("/books/:id", controllers.DeleteBook)

	// Adding book routes
	bookRoute := r.Group("/books")
	bookRoute.Use(middlewares.JwtAuthMiddleware())   // Protect book routes with JWT Middleware
	bookRoute.GET("", controllers.GetAllBooks)       // Get all books
	bookRoute.GET("/:id", controllers.GetBookByID)   // Get a single book by ID
	bookRoute.POST("", controllers.CreateBook)       // Create a new book
	bookRoute.PATCH("/:id", controllers.UpdateBook)  // Update an existing book
	bookRoute.DELETE("/:id", controllers.DeleteBook) // Delete a book

	// Di dalam fungsi SetupRouter

	// Routing Mahasiswa
	r.GET("/mahasiswa", controllers.GetAllMahasiswa)
	r.POST("/mahasiswa", controllers.CreateMahasiswa)
	r.GET("/mahasiswa/:id", controllers.GetMahasiswaByID)
	r.PUT("/mahasiswa/:id", controllers.UpdateMahasiswa)
	r.DELETE("/mahasiswa/:id", controllers.DeleteMahasiswa)

	// Routing Dosen
	r.GET("/dosen", controllers.GetAllDosen)
	r.POST("/dosen", controllers.CreateDosen)
	r.GET("/dosen/:id", controllers.GetDosenByID)
	r.PUT("/dosen/:id", controllers.UpdateDosen)
	r.DELETE("/dosen/:id", controllers.DeleteDosen)

	// Routing Mata Kuliah
	r.GET("/mata-kuliah", controllers.GetAllMataKuliah)
	r.POST("/mata-kuliah", controllers.CreateMataKuliah)
	r.GET("/mata-kuliah/:id", controllers.GetMataKuliahByID)
	r.PUT("/mata-kuliah/:id", controllers.UpdateMataKuliah)
	r.DELETE("/mata-kuliah/:id", controllers.DeleteMataKuliah)

	// Routing Jadwal Kuliah
	r.GET("/jadwal-kuliah", controllers.GetAllJadwalKuliah)
	r.POST("/jadwal-kuliah", controllers.CreateJadwalKuliah)
	r.GET("/jadwal-kuliah/:id", controllers.GetJadwalKuliahByID)
	r.PUT("/jadwal-kuliah/:id", controllers.UpdateJadwalKuliah)
	r.DELETE("/jadwal-kuliah/:id", controllers.DeleteJadwalKuliah)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
