package main

import (
	"fmt"
	"os"

	cfg "github.com/ozonva/ova-book-api/internals/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Не передан путь до файла конфигурации.")
		return
	}

	configPath := os.Args[1]
	cfg.ReadConfig(configPath)
}
