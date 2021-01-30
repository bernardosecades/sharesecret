package main

import (
	"fmt"
	"github.com/bernardosecades/sharesecret/repository"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var (
	commitHash string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Not .env file found")
	}
}

func main() {

	fmt.Printf("Build Time: %s\n", time.Now().Format(time.RFC3339))
	fmt.Printf("Version: %s\n", commitHash)

	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	secretRepository := repository.NewMySqlSecretRepository(dbName, dbUser, dbPass, dbHost, dbPort)
	r, err := secretRepository.RemoveSecretsExpired()

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	fmt.Printf("%d secrets deleted", r)
}