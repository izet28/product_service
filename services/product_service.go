package services

import (
	"product_service/models"
	"product_service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

// ProductService mendefinisikan logika bisnis untuk produk
type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(id uint) error
}

// productService adalah implementasi dari ProductService
type productService struct {
	repository repository.ProductRepository
	validate   *validator.Validate
	logger     *logrus.Logger
}

// NewProductService membuat instance baru dari ProductService
func NewProductService(repository repository.ProductRepository, logger *logrus.Logger) ProductService {
	return &productService{
		repository: repository,
		validate:   validator.New(),
		logger:     logger,
	}
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	s.logger.Info("Mengambil semua produk")
	return s.repository.FindAll()
}

func (s *productService) GetProductByID(id uint) (models.Product, error) {
	s.logger.Infof("Mengambil produk dengan ID: %d", id)
	return s.repository.FindByID(id)
}

func (s *productService) CreateProduct(product models.Product) (models.Product, error) {
	s.logger.Infof("Membuat produk baru: %s", product.Name)
	// Validasi input
	err := s.validate.Struct(product)
	if err != nil {
		s.logger.Error("Validasi gagal: ", err)
		return models.Product{}, err
	}

	return s.repository.Create(product)
}

func (s *productService) UpdateProduct(product models.Product) (models.Product, error) {
	s.logger.Infof("Memperbarui produk dengan ID: %d", product.ID)
	// Validasi input
	err := s.validate.Struct(product)
	if err != nil {
		s.logger.Error("Validasi gagal: ", err)
		return models.Product{}, err
	}

	return s.repository.Update(product)
}

func (s *productService) DeleteProduct(id uint) error {
	s.logger.Infof("Menghapus produk dengan ID: %d", id)
	product, err := s.repository.FindByID(id)
	if err != nil {
		s.logger.Error("Produk tidak ditemukan: ", err)
		return err
	}
	return s.repository.Delete(product)
}
