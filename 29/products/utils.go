package products

func GetProduct() *Product {
	return &Product{
		Id:    "1",
		Name:  "product1",
		Price: 123,
	}
}
