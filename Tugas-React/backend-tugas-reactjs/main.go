// main.go
package main

import (
	"backend-tugas-reactjs/config"
	"backend-tugas-reactjs/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDataBase() // Pastikan ini sesuai dengan fungsi setup DB Anda
	r := gin.Default()
	routes.SetupRouter(db, r) // Ini menghubungkan router Anda

	r.Run() //
}
