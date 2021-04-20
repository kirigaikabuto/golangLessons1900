package users

import (
	"database/sql"
	"github.com/kirigaikabuto/golangLessons1900/17/config"
	_ "github.com/lib/pq"
	"log"
)

var Queries = []string{
	`CREATE TABLE IF NOT EXISTS users (
		id serial,
		username text,
		password text,
		PRIMARY KEY(id)
	);`,
}

type usersPostgreStore struct {
	db *sql.DB
}

func NewUsersPostgreStore(configPostgre config.Config) (UsersStore, error) {
	db, err := config.GetDbConn(config.GetConnString(configPostgre))
	if err != nil {
		return nil, err
	}
	for _, q := range Queries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	return &usersPostgreStore{db: db}, nil
}

func (u *usersPostgreStore) AddUser(user User) error {
	err := u.db.QueryRow("insert into users (username,password) values ($1,$2) returning id", user.Username, user.Password).Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *usersPostgreStore) GetById(id int) (*User, error) {
	product := &User{}
	err := u.db.QueryRow("select username,password from users where id= $1", id).Scan(&product.Username, &product.Password)
	if err != nil {
		return nil, err
	}
	product.Id = id
	return product, nil
}
