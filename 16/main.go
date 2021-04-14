package main

import (
	"errors"
	"fmt"
	"github.com/kirigaikabuto/golangLessons1900/16/config"
	"github.com/kirigaikabuto/golangLessons1900/16/users"
	"github.com/urfave/cli"
	"log"
	"os"
)

var (
	id            = 0
	usersStore    *users.UsersStore
	err           error
	username      string
	password      string
	loginCmdFlags = []cli.Flag{
		cli.StringFlag{
			Name:        "username,u",
			Value:       "",
			Destination: &username,
			Usage:       "flag for setting username",
		},
		cli.StringFlag{
			Name:        "password,p",
			Value:       "",
			Destination: &password,
			Usage:       "flag for setting password",
		},
	}
)

func loginAction(c *cli.Context) error {
	if username == "" || password == "" {
		return errors.New("please set values to username and password")
	}
	user := users.User{
		Username: username,
		Password: password,
	}
	usersData, err := usersStore.List()
	if err != nil {
		return err
	}
	isExist := false
	for _, v := range usersData {
		if v.Username == user.Username && v.Password == v.Password {
			user = v
			isExist = true
		}
	}
	if isExist {
		id = user.Id
		fmt.Printf("Welcome %d", user.Id)
	} else {
		log.Fatal(errors.New("no user by this username and password"))
	}
	return nil
}

func main() {
	cf := config.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "testdb",
		CollectionName: "users",
	}
	usersStore, err = users.NewUsersStore(cf)
	if err != nil {
		log.Fatal(err)
	}
	app := cli.NewApp()
	app.Name = "App for user-management"
	app.Description = "App for user-management"
	app.Usage = "first of all please login"
	app.Commands = []cli.Command{
		{
			Name:    "login",
			Aliases: []string{"log", "l"},
			Usage:   "login to the system",
			Action:  loginAction,
			Flags:   loginCmdFlags,
		},
	}
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
