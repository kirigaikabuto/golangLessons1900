package users

type UsersStore interface {
	AddUser(user User) error
	GetById(id int) (*User, error)
	//List() ([]User, error)
	//Update(user User) error
	//Delete(id int) error
}
