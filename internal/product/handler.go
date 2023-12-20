package product

import (
	"context"
	"net/http"

	"github.com/tinchourteaga/go-grpc-product-svc/internal/models"
	"github.com/tinchourteaga/go-grpc-product-svc/internal/pb"
)

type Product struct {
	productSvc Service
}

func NewProduct(svc Service) *Product {
	return &Product{
		productSvc: svc,
	}
}

func (p *Product) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	product := models.Product{Name: req.Name, Stock: req.Stock, Price: req.Price}
	id, err := p.productSvc.Create(ctx, product)
	if err != nil {
		return &pb.CreateResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &pb.CreateResponse{
		Status: http.StatusCreated,
		Id:     id,
	}, nil
}

func (p *Product) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	product, err := p.productSvc.FindOne(ctx, req.Id)
	if err != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:    product.Id,
		Name:  product.Name,
		Stock: product.Stock,
		Price: product.Price,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (p *Product) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	stkDecreaseLog := models.StockDecreaseLog{Id: req.Id, OrderId: req.OrderId}
	err := p.productSvc.DecreaseStock(ctx, stkDecreaseLog)
	if err != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}
