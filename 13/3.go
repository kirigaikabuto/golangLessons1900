package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

//глобальная переменная которая доступна отовсюду
var (
	znak    string
	number1 int
	number2 int
)

//go mod init
//go get github.com/urfave/cli
//go run 2.go -> hello from cli app
//go run 2.go help
func MainAction2(c *cli.Context) error {
	total := 0
	if znak == "+" {
		total = number1 + number2
	} else if znak == "-" {
		total = number1 - number2
	}
	fmt.Println(total)
	return nil
}

func main() {
	//создание приложения
	app := cli.NewApp()
	//указываем свойства
	app.Name = "App for downloading video"
	app.Description = "Just set the url and video files will bee downloaded"
	app.Usage = "Set parameter url to link"
	//создаем флаги
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "znak,z",
			Usage:       "flag for containing string value",
			Value:       "+",
			Destination: &znak,
		},
		cli.IntFlag{
			Name:        "number1,n1",
			Usage:       "flag for containing int value",
			Value:       0,
			Destination: &number1,
		},
		cli.IntFlag{
			Name:        "number2,n2",
			Usage:       "flag for containing int value",
			Value:       0,
			Destination: &number2,
		},
	}
	//функция которая срабатывает автоматически когда запускаем приложение
	//c *cli.Context -> в ней хранятся все аргументы и флаги которые указываются при запуске приложения
	app.Action = MainAction2
	//запуск приложения
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
