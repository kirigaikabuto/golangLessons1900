package movie

type MovieStore interface {
	Create(movie *Movie) (*Movie, error)
	Delete(id int) error
	List() ([]Movie, error)
	Get(id int) (*Movie, error)
}
