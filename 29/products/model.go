package products

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type GetProductByIdCommand struct {
	Id string `json:"id"`
}
