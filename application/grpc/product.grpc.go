package grpc

import (
	"context"

	"github.com/VictorMayer/first-grpc-server/application/grpc/pb"
	"github.com/VictorMayer/first-grpc-server/application/usecase"
)

type ProductGrpcService struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func (p *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := p.ProductUseCase.CreateProduct(in.Name, in.Description, in.Price)
	if err != nil {
		return &pb.CreateProductResponse{
			Product: &pb.Product{},
		}, err
	}
	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		},
	}, nil
}

func (p *ProductGrpcService) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	product, err := p.ProductUseCase.FindAllProducts()
	if err != nil {
		return &pb.FindProductsResponse{}, err
	}
	return &pb.FindProductsResponse{
		Products: []*pb.Product{
			{
				Name:        product[0].Name,
				Description: product[0].Description,
				Price:       product[0].Price,
			},
		},
	}, nil
}

func NewProductGrpcService(usecase usecase.ProductUseCase) *ProductGrpcService {
	return &ProductGrpcService{
		ProductUseCase: usecase,
	}
}
