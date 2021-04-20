package products

type ProductStore interface {
	CreateProduct(product *Product) (*Product, error)
	GetProduct(id int64) (*Product, error)
	GetProduct123(id int64) (*Product, error)
	//delete/list/update mongo
}
