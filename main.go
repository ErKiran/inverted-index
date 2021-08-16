package main

import (
	"fmt"
	"fts/cmd"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("err", err)
		os.Exit(0)
	}
	cmd.Execute()
}
