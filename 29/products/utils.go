package products

func GetProduct() *Product {
	return &Product{
		Id:    "1",
		Name:  "product1",
		Price: 123,
	}
}

func GetProducts() []Product {
	products := []Product{
		{
			Id:    "1",
			Name:  "product1",
			Price: 123,
		},
		{
			Id:    "2",
			Name:  "product2",
			Price: 456,
		},
		{
			Id:    "3",
			Name:  "product3",
			Price: 179,
		},
	}
	return products
}

func GetProductById(id string) *Product {
	products := []Product{
		{
			Id:    "1",
			Name:  "product1",
			Price: 123,
		},
		{
			Id:    "2",
			Name:  "product2",
			Price: 456,
		},
		{
			Id:    "3",
			Name:  "product3",
			Price: 179,
		},
	}
	for _, v := range products {
		if v.Id == id {
			return &v
		}
	}
	return &Product{}
}
