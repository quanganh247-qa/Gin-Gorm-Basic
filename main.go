package main

import (
	"log"

	"github.com/quanganh247-qa/gorm-project/app/server"
	"github.com/quanganh247-qa/gorm-project/app/util"
)

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
