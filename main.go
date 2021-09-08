package main

import (
	"log"
	"os"

	"github.com/jtheo/milestone1-code/storage"
	"github.com/jtheo/milestone1-code/web"
)

func main() {
	file := os.Getenv("DATA_FILE_PATH")
	if file == "" {
		log.Fatal("Missing variable DATA_FILE_PATH")
	}

	ids := storage.Storage{}
	err := ids.InitMapFile(file)
	if err != nil {
		log.Fatalf("Error reading file %s: %v\n", file, err)
	}

	web.Run(ids)
}
