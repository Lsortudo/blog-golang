package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/lsortudo/blog-golang/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DBDriver, DBUser, DBPassword, DBPort, DBHost, DBName string) {
	var err error
	// Posso adicionar outro banco como por exemplo o MYSQL pra caso suba aplicacao fazendo dois bancos, local e la
	if DBDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)
		server.DB, err = gorm.Open(DBDriver, DBURL)
		if err != nil {
			fmt.Printf("Nao consigo me conectar ao banco de dados: %s", DBDriver)
			log.Fatal("Confira seu erro -> ", err)
		} else {
			fmt.Printf("Estamos conectados ao banco de dados %s com sucesso!!!", DBDriver)
		}
	}
	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) // Migrar
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Porta 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
