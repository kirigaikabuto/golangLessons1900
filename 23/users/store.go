package users

type UsersStore interface {
	Create(user *User) (*User, error)
	Get(id int) (*User, error)
	GetByUsernameAndPassword(username, password string) (*User, error)
}
