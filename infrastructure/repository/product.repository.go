package repository

import (
	"fmt"

	"github.com/VictorMayer/first-grcp-server/domain/model"
	"github.com/jinzhu/gorm"
)

// type ProductRepositoryInterface interface {
// 	CreateProduct(product *Product) (*Product, error)
// 	FindProductById(id string) (*Product, error)
// 	FindProductsByName(name string) ([]*Product, error)
// 	FindAllProducts() ([]*Product, error)
// }

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func (r *ProductRepositoryDb) AddProduct(product *model.Product) error {
	err := r.Db.Create(product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryDb) FindProductById(id string) (*model.Product, error) {
	var product model.Product
	r.Db.First(&product, "id = ?", id)

	if product.ID == "" {
		return nil, fmt.Errorf("product not found")
	}
	return &product, nil
}

func (r *ProductRepositoryDb) FindProductsByName(name string) (*model.Product, error) {
	var product model.Product
	r.Db.Where("name %s", name).Find(&product)

	if product.ID == "" {
		return nil, fmt.Errorf("product not found")
	}
	return &product, nil
}

func (r *ProductRepositoryDb) FindAllProducts() ([]*model.Product, error) {
	var products []*model.Product
	r.Db.Find(&products)

	if len(products) == 0 {
		return nil, fmt.Errorf("no products found")
	}
	return products, nil
}
