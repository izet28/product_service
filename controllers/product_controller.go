package controllers

import (
	"encoding/json"
	"net/http"
	"product_service/models"
	"product_service/services"
	"product_service/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// ProductController menangani permintaan HTTP untuk produk
type ProductController struct {
	service services.ProductService
	logger  *logrus.Logger
}

// NewProductController membuat instance baru dari ProductController
func NewProductController(service services.ProductService, logger *logrus.Logger) *ProductController {
	return &ProductController{
		service: service,
		logger:  logger,
	}
}

// GetAllProducts mengembalikan daftar semua produk
func (c *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.service.GetAllProducts()
	if err != nil {
		c.logger.Error("Gagal mengambil produk: ", err)
		utils.RespondError(w, http.StatusInternalServerError, "Gagal mengambil produk")
		return
	}
	utils.RespondJSON(w, http.StatusOK, products)
}

// GetProductByID mengembalikan detail produk berdasarkan ID
func (c *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		c.logger.Warn("ID produk tidak valid: ", params["id"])
		utils.RespondError(w, http.StatusBadRequest, "ID produk tidak valid")
		return
	}
	product, err := c.service.GetProductByID(uint(id))
	if err != nil {
		c.logger.Error("Produk tidak ditemukan: ", err)
		utils.RespondError(w, http.StatusNotFound, "Produk tidak ditemukan")
		return
	}
	utils.RespondJSON(w, http.StatusOK, product)
}

// CreateProduct membuat produk baru
func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		c.logger.Warn("Input tidak valid: ", err)
		utils.RespondError(w, http.StatusBadRequest, "Input tidak valid")
		return
	}
	createdProduct, err := c.service.CreateProduct(product)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			c.logger.Error("Validasi input gagal: ", err)
			utils.RespondError(w, http.StatusBadRequest, "Validasi input gagal")
		} else {
			c.logger.Error("Gagal membuat produk: ", err)
			utils.RespondError(w, http.StatusInternalServerError, "Gagal membuat produk")
		}
		return
	}
	utils.RespondJSON(w, http.StatusCreated, createdProduct)
}

// UpdateProduct memperbarui produk yang ada
func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		c.logger.Warn("ID produk tidak valid: ", params["id"])
		utils.RespondError(w, http.StatusBadRequest, "ID produk tidak valid")
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		c.logger.Warn("Input tidak valid: ", err)
		utils.RespondError(w, http.StatusBadRequest, "Input tidak valid")
		return
	}
	product.ID = uint(id)
	updatedProduct, err := c.service.UpdateProduct(product)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			c.logger.Warn("Validasi input gagal: ", err)
			utils.RespondError(w, http.StatusBadRequest, "Validasi input gagal")
		} else {
			c.logger.Error("Gagal memperbarui produk: ", err)
			utils.RespondError(w, http.StatusInternalServerError, "Gagal memperbarui produk")
		}
		return
	}
	utils.RespondJSON(w, http.StatusOK, updatedProduct)
}

// DeleteProduct menghapus produk berdasarkan ID
func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		c.logger.Warn("ID produk tidak valid: ", params["id"])
		utils.RespondError(w, http.StatusBadRequest, "ID produk tidak valid")
		return
	}
	err = c.service.DeleteProduct(uint(id))
	if err != nil {
		c.logger.Error("Gagal menghapus produk: ", err)
		utils.RespondError(w, http.StatusInternalServerError, "Gagal menghapus produk")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
