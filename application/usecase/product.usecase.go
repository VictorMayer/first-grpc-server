package usecase

import "github.com/VictorMayer/first-grcp-server/domain/model"

type ProductUseCase struct {
	ProductRepository model.ProductRepositoryInterface
}

func (p *ProductUseCase) AddProduct(name string, description string, price float64) (*model.Product, error) {
	product, err := model.NewProduct(name, description, price)
	if err != nil {
		return nil, err
	}

	p.ProductRepository.CreateProduct(product)
	if product.ID == "" {
		return nil, err
	}
	return product, nil
}

func (p *ProductUseCase) FindProductById(id string) (*model.Product, error) {
	product, err := p.ProductRepository.FindProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductUseCase) FindProductsByName(name string) ([]*model.Product, error) {
	products, err := p.ProductRepository.FindProductsByName(name)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductUseCase) FindAllProducts() ([]*model.Product, error) {
	products, err := p.ProductRepository.FindAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}
