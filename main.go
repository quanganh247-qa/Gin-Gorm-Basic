package main

import (
	"log"

	_ "github.com/quanganh247-qa/gorm-project/docs"

	"github.com/quanganh247-qa/gorm-project/app/server"
	"github.com/quanganh247-qa/gorm-project/app/util"
)

// @title 	Notes Service API
// @version	1.0
// @description APIs in Go using Gin framework

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @host 	localhost:8080
// @BasePath /api/v1
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load config: ", err)
	}

	server := runGinServer(*config)
	defer func() {
		server.Connection.Close()
	}()

	// routerDefault.Run(":8080")

}

func runGinServer(config util.Config) *server.Server {
	server, err := server.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	return server
}
