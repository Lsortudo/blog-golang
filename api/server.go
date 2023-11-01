package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lsortudo/blog-golang/api/controllers"
	"github.com/lsortudo/blog-golang/api/seedDB"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Erro no ENV %v", err)
	} else {
		fmt.Println("Env values corretos")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seeddb.Load(server.DB)

	server.Run(":8080")
}
