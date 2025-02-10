package creational

type Product interface {
	Use() string
}

type ContractProduct struct{}

func (p *ContractProduct) Use() string {
	return "Using contract product"
}

type Creator interface {
	CreateProduct() Product
}

type ConcreteCreator struct{}

func (c *ConcreteCreator) CreateProduct() Product {
	return &ContractProduct{}
}
