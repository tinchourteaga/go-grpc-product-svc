package product

import (
	"context"

	"github.com/tinchourteaga/go-grpc-product-svc/internal/models"
)

type Service interface {
	Create(context.Context, models.Product) (int64, error)
	FindOne(context.Context, int64) (models.Product, error)
	DecreaseStock(context.Context, models.StockDecreaseLog) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) Create(ctx context.Context, product models.Product) (int64, error) {
	id, err := s.repository.Create(ctx, product)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *service) FindOne(ctx context.Context, id int64) (models.Product, error) {
	product, err := s.repository.FindOne(ctx, id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (s *service) DecreaseStock(ctx context.Context, stkDecreaseLog models.StockDecreaseLog) error {
	err := s.repository.DecreaseStock(ctx, stkDecreaseLog)
	if err != nil {
		return err
	}
	return nil
}
