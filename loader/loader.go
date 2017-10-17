package loader

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadVars() string {
	err := godotenv.Load()

	if err != nil {
		cwd, _ := os.Getwd()
		fmt.Println("CWD: ", cwd)
		log.Fatal("Error loading .env file in path: ", err)
	}

	return os.Getenv("SOME_VAR")
}
