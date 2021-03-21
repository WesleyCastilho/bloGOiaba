package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/WesleyCastilho/blog/api/controllers"
	//"github.com/WesleyCastilho/blog/api/seed"
)

var server = controllers.Server{}

func Run() {

	var err error
	var port = os.Getenv("PORT")
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("Setando variaveis de ambiente")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	//seed.Load(server.DB)

	server.Run(port)

}