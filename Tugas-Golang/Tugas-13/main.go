package main

import (
	"Tugas-13/controllers" // Adjust this import to match the actual path to your controllers package
	"Tugas-13/utils"       // Adjust this import to match the actual path to your utils package
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// Routes for Mahasiswa
	router.GET("/mahasiswa", controllers.GetAllMahasiswa)
	router.GET("/mahasiswa/:id", controllers.GetMahasiswaByID)
	router.POST("/mahasiswa", utils.BasicAuth(utils.HttpRouterToHandlerFunc(controllers.CreateMahasiswa)))
	router.PUT("/mahasiswa/:id", utils.BasicAuth(utils.HttpRouterToHandlerFunc(controllers.UpdateMahasiswa)))
	router.DELETE("/mahasiswa/:id", utils.BasicAuth(utils.HttpRouterToHandlerFunc(controllers.DeleteMahasiswa)))

	// Routes for Nilai
	router.GET("/nilai", controllers.GetAllNilai)
	router.POST("/nilai", utils.BasicAuth(utils.HttpRouterToHandlerFunc(controllers.CreateNilai)))
	router.PUT("/nilai/:id", utils.BasicAuth(utils.HttpRouterToHandlerFunc(controllers.UpdateNilai)))
	router.DELETE("/nilai/:id", utils.BasicAuth(utils.HttpRouterToHandlerFunc(controllers.DeleteNilai)))
	router.GET("/nilai/mahasiswa/:id", controllers.GetNilaiByMahasiswaID)
	// router.GET("/nilai/:id", controllers.GetNilaiByID)

	// Routes for Mata Kuliah
	router.GET("/matkul", controllers.GetAllMataKuliah)
	router.GET("/matkul/:id", controllers.GetMataKuliahByID)
	router.POST("/matkul", utils.BasicAuth(utils.HttpRouterToHandlerFunc(controllers.CreateMataKuliah)))
	router.PUT("/matkul/:id", utils.BasicAuth(utils.HttpRouterToHandlerFunc(controllers.UpdateMataKuliah)))
	router.DELETE("/matkul/:id", utils.BasicAuth(utils.HttpRouterToHandlerFunc(controllers.DeleteMataKuliah)))

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
