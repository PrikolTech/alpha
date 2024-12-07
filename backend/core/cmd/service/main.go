package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	os.Exit(run())
}

func run() int {
	godotenv.Overload()
	return 0
}
