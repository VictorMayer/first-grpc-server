package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type ProductRepositoryInterface interface {
	CreateProduct(product *Product) (*Product, error)
	FindProductById(id string) (*Product, error)
	FindProductsByName(name string) ([]*Product, error)
	FindAllProducts() ([]*Product, error)
}

type Product struct {
	Base        `valid:"required"`
	Name        string  `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Description string  `json:"description" gorm:"type:varchar(255)"`
	Price       float64 `json:"price" gorm:"type:float" valid:"notnull"`
}

func (product *Product) isValid() error {
	_, err := govalidator.ValidateStruct(product)

	if err != nil {
		return err
	}
	return nil
}

func NewProduct(name string, description string, price float64) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	product.ID = uuid.NewV4().String()
	product.CreatedAt = time.Now()

	error := product.isValid()
	if error != nil {
		return nil, error
	}
	return &product, nil
}
