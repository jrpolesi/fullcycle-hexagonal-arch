package dto

import "github.com/jrpolesi/fullcycle-hexagonal-arch/application"

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) ToProduct(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}
	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}

	return product, nil
}

func (p *Product) ToDTO(product application.ProductInterface) *Product {
	if product.GetID() != "" {
		p.ID = product.GetID()
	}
	p.Name = product.GetName()
	p.Price = product.GetPrice()
	p.Status = product.GetStatus()

	return p
}
