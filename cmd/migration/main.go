package main

import (
	"flag"
	"fmt"
	_ "go-gin/migrations"
	"go-gin/svc/sqlx"
	"go-gin/utils/filex"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/color"

	"github.com/pressly/goose"
)

var (
	configFile = flag.String("f", "./.env.yaml", "the config file")
	dir        = "./migrations"
	flags      = flag.NewFlagSet("goose", flag.ExitOnError)
)

type clog struct {
}

var _ goose.Logger = &clog{}

func (*clog) Fatal(v ...interface{}) {
	color.Println(color.Red(fmt.Sprint(v...)))
}

func (*clog) Fatalf(format string, v ...interface{}) {
	color.Println(color.Red(fmt.Sprintf(format, v...)))
}

func (*clog) Print(v ...interface{}) {
	color.Println(color.White(fmt.Sprint(v...)))

}
func (*clog) Println(v ...interface{}) {
	color.Println(color.White(fmt.Sprint(v...)))
}
func (*clog) Printf(format string, v ...interface{}) {
	color.Println(color.Green(fmt.Sprintf(format, v...)))
}

type Config struct {
	Mysql sqlx.Config `yaml:"Mysql"`
}

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()
	color.Enable()
	goose.SetLogger(&clog{})

	if len(args) < 1 {
		color.Println(color.Red(`please input the execute command[up|down|create]`))
		return
	}

	command := args[0]

	var c Config
	filex.MustLoad(*configFile, &c)

	db, err := goose.OpenDBWithDriver("mysql", c.Mysql.DataSource)
	if err != nil {
		color.Println(color.Red(fmt.Sprintf("goose: failed to open DB: %v\n", err)))
	}
	defer func() {
		if err := db.Close(); err != nil {
			color.Println(color.Red(fmt.Sprintf("goose: failed to close DB: %v\n", err)))
		}
	}()

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}
	goose.SetTableName("migrations")

	if err := goose.Run(command, db, dir, arguments...); err != nil {
		color.Println(color.Red(fmt.Sprintf("goose %v: %v", command, err)))
	}
}
