package products

type ProductStore interface {
	CreateProduct(product *Product) (*Product, error)
	GetProduct(id int64) (*Product, error)
}
