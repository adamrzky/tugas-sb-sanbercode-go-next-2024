package main

import (
	"final-project-golang-individu/config"
	_ "final-project-golang-individu/docs" // Import generated docs
	"final-project-golang-individu/middleware"
	"final-project-golang-individu/routes"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Final Project Golang API
// @version 1.0
// @description This is a sample server for a culinary review app.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	config.InitDB()

	r := gin.Default()

	r.Use(middleware.CORS())

	routes.Setup(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
