package main

import (
	"flag"
	"fmt"
	_ "go-gin/migrations"
	"go-gin/svc/sqlx"
	"go-gin/utils/filex"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/pressly/goose"
)

var (
	configFile = flag.String("f", "./.env.yaml", "the config file")
	dir        = "./migrations"
	flags      = flag.NewFlagSet("goose", flag.ExitOnError)
)

type Config struct {
	Mysql sqlx.Config `yaml:"Mysql"`
}

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()
	fmt.Println(args)
	if len(args) < 1 {
		log.Fatal("please input the execute command[up|down|create]\n")
		return
	}

	command := args[0]

	var c Config
	filex.MustLoad(*configFile, &c)

	db, err := goose.OpenDBWithDriver("mysql", c.Mysql.DataSource)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}
	goose.SetTableName("migrations")
	if err := goose.Run(command, db, dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
