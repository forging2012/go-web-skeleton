package main

import (
	"fmt"
	// boot and init some services(log, cache, eureka)
	"github.com/inhere/go-web-skeleton/app"

	// init redis, mongo, mysql connection
	_ "github.com/inhere/go-web-skeleton/model/mongo"
	_ "github.com/inhere/go-web-skeleton/model/mysql"
	_ "github.com/inhere/go-web-skeleton/model/rds"

	"log"
	"os"
	"github.com/gookit/sux/handlers"
	"github.com/gookit/sux"
)

func init() {
	app.Boot()
}

func main() {
	r := sux.New()

	if app.IsEnv(app.DEV) {
		sux.Debug(true)
	}

	// global middleware
	r.Use(handlers.RequestLogger())

	app.AddRoutes(r)

	log.Printf("======================== Begin Running(PID: %d) ========================", os.Getpid())

	// default is listen and serve on 0.0.0.0:8080
	r.Listen(fmt.Sprintf("0.0.0.0:%d", app.HttpPort))
}
