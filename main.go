package main

import (
	"github.com/mmattklaus/go-jwt-demo/config"
	"github.com/mmattklaus/go-jwt-demo/database"
	"github.com/mmattklaus/go-jwt-demo/models"
	"github.com/mmattklaus/go-jwt-demo/router"
	"log"
	"net/http"
	"os"
)

var dh database.Database
var logger *log.Logger
var conf config.Config

var (
	User models.User
)

func main() {

	logger = log.New(os.Stdout, "Micros: ", log.LstdFlags|log.Lshortfile)
	conf.Read()
	dh.Connect(&conf, logger)
	defer dh.DB.Close()

	//dh.DB.AutoMigrate()


	api := router.NewAPI(dh.DB, &conf, logger)
	api.InitializeRoutes()

	logger.Println("server started on " + conf.ServerAddr)

	// svr := server.New(mux, conf.ServerAddr)
	logger.Fatalln(http.ListenAndServe(conf.ServerAddr, nil))
	//logger.Println("Hello, World!")
	//if err != nil {
	//	logger.Fatalf("Error starting server: %v", err.Error())
	//}


}
