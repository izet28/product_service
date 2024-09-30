package main

import (
	"log"
	"net/http"
	"product_service/config"
	"product_service/controllers"
	"product_service/repository"
	"product_service/services"
	"product_service/utils"

	"github.com/gorilla/mux"
)

func main() {
	// Inisialisasi logger
	logger := utils.InitLogger()

	// Inisialisasi konfigurasi dan koneksi ke database
	db := config.InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Inisialisasi repository, service, dan controller
	productRepository := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepository, logger)
	productController := controllers.NewProductController(productService, logger)

	// Buat router baru
	r := mux.NewRouter()

	// Daftarkan endpoint ke router
	r.HandleFunc("/products", productController.GetAllProducts).Methods("GET")
	r.HandleFunc("/products/{id}", productController.GetProductByID).Methods("GET")
	r.HandleFunc("/products", productController.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", productController.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", productController.DeleteProduct).Methods("DELETE")

	// Jalankan server
	log.Println("Server berjalan pada port 8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
