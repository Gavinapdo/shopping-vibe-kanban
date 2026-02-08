package product

import (
	"errors"
	"strings"
)

var (
	ErrProductNotFound = errors.New("商品不存在")
	ErrInvalidInput    = errors.New("商品参数不合法")
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListProducts() []Product {
	return s.repo.List()
}

func (s *Service) GetProduct(id int64) (Product, error) {
	p, ok := s.repo.GetByID(id)
	if !ok {
		return Product{}, ErrProductNotFound
	}
	return p, nil
}

func (s *Service) CreateProduct(input CreateInput) (Product, error) {
	if strings.TrimSpace(input.Name) == "" || input.Price <= 0 || input.Stock < 0 {
		return Product{}, ErrInvalidInput
	}

	return s.repo.Create(Product{
		Name:        strings.TrimSpace(input.Name),
		Description: strings.TrimSpace(input.Description),
		Price:       input.Price,
		Stock:       input.Stock,
	}), nil
}

func (s *Service) UpdateProduct(id int64, input UpdateInput) (Product, error) {
	if strings.TrimSpace(input.Name) == "" || input.Price <= 0 || input.Stock < 0 {
		return Product{}, ErrInvalidInput
	}

	updated, ok := s.repo.Update(id, Product{
		Name:        strings.TrimSpace(input.Name),
		Description: strings.TrimSpace(input.Description),
		Price:       input.Price,
		Stock:       input.Stock,
	})
	if !ok {
		return Product{}, ErrProductNotFound
	}

	return updated, nil
}

func (s *Service) DeleteProduct(id int64) error {
	if ok := s.repo.Delete(id); !ok {
		return ErrProductNotFound
	}
	return nil
}
