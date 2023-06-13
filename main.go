package main

import (
	"fmt"
	variables "go-basics/01-variables"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))

	// modules.Main()   //* 00-modules

	variables.Main() //* 01-variables
}
