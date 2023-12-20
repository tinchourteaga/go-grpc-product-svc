package product

import (
	"context"
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/tinchourteaga/go-grpc-product-svc/internal/models"
	"github.com/tinchourteaga/go-grpc-product-svc/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	Create(context.Context, models.Product) (int64, error)
	FindOne(context.Context, int64) (models.Product, error)
	DecreaseStock(context.Context, models.StockDecreaseLog) error
	IncreaseStock(context.Context, models.Product) error
	Exists(context.Context, models.Product) bool
}

type repository struct {
	con db.Connector
}

func NewRepository(con db.Connector) Repository {
	return &repository{
		con: con,
	}
}

func (r *repository) Create(ctx context.Context, product models.Product) (int64, error) {
	exists := r.Exists(ctx, product)
	if exists {
		err := r.IncreaseStock(ctx, product)
		if err != nil {
			return 0, err
		}
		return 0, nil
	}

	r.con.DB.Create(&product)
	return product.Id, nil
}

func (r *repository) FindOne(ctx context.Context, id int64) (models.Product, error) {
	product := models.Product{}
	result := r.con.DB.First(&product, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Product{}, gorm.ErrRecordNotFound
	}

	return product, nil
}

func (r *repository) DecreaseStock(ctx context.Context, stkDecreaseLog models.StockDecreaseLog) error {
	dbProduct := models.Product{}
	result := r.con.DB.First(&dbProduct, stkDecreaseLog.Id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	if dbProduct.Stock < 1 {
		return errors.New("no stock for product: " + dbProduct.Name)
	}

	dbStkDecreaseLog := models.StockDecreaseLog{}
	result = r.con.DB.Where(&models.StockDecreaseLog{OrderId: stkDecreaseLog.OrderId}).First(&dbStkDecreaseLog)
	if result.Error == nil {
		return errors.New("stock already decreased")
	}

	dbProduct.Stock = dbProduct.Stock - 1
	r.con.DB.Save(&dbProduct)

	dbStkDecreaseLog.OrderId = stkDecreaseLog.OrderId
	dbStkDecreaseLog.ProductRefer = dbProduct.Id

	r.con.DB.Create(&dbStkDecreaseLog)

	log.Info().Msg("stock decreased for product " + dbProduct.Name)
	return nil
}

func (r *repository) IncreaseStock(ctx context.Context, product models.Product) error {
	dbProduct := models.Product{}
	result := r.con.DB.Where(&models.Product{Name: product.Name, Price: product.Price}).First(&dbProduct)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	dbProduct.Stock = dbProduct.Stock + product.Stock
	r.con.DB.Save(&dbProduct)

	log.Info().Msg("stock increased by " + strconv.Itoa(int(product.Stock)) + " for product " + product.Name)
	return nil
}

func (r *repository) Exists(ctx context.Context, product models.Product) bool {
	result := r.con.DB.Where(&models.Product{Name: product.Name, Price: product.Price}).First(&models.Product{})
	return result.Error == nil
}
