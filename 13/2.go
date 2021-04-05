package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"strconv"
)

//go mod init
//go get github.com/urfave/cli
//go run 2.go -> hello from cli app
//go run 2.go help
func MainAction(c *cli.Context) error {
	//var s string
	//fmt.Scanf("%s", &s)
	//fmt.Println(s)
	a := c.Args().Get(0)
	b := c.Args().Get(1)
	if c.NArg() == 2 {
		num1, err := strconv.Atoi(a)
		if err != nil {
			return err
		}
		num2, err := strconv.Atoi(b)
		if err != nil {
			return err
		}
		fmt.Println(num1 + num2)
	} else {
		return errors.New("count of arguments should be only 2")
	}

	return nil
}

func main() {
	//создание приложения
	app := cli.NewApp()
	//указываем свойства
	app.Name = "App for downloading video"
	app.Description = "Just set the url and video files will bee downloaded"
	app.Usage = "Set parameter url to link"
	//функция которая срабатывает автоматически когда запускаем приложение
	//c *cli.Context -> в ней хранятся все аргументы и флаги которые указываются при запуске приложения
	app.Action = MainAction
	//запуск приложения
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
