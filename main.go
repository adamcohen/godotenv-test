package main

import (
	"fmt"

	"github.com/adamcohen/godotenv-test/loader"
)

func main() {
	value := loader.LoadVars()
	fmt.Println("VALUE: ", value)
}
