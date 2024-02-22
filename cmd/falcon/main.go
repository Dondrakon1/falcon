package main

import (
	"falcon/internal/camera"
	"falcon/internal/service/code"
	"falcon/internal/storage/sqlite"
	"log"
)

func main() {
	log.Println("Initializing Falcon...")
	log.Println("Loading storage...")
	storage, err := sqlite.New("./storage/falcon.db")
	if err != nil {
		panic(err)
	}
	log.Println("Loading service...")
	srv := code.NewCodeService(storage)
	log.Println("Loading camera...")
	cam, err := camera.NewCamera("192.168.252.37:2003", srv)
	if err != nil {
		panic(err)

	}
	log.Println("Starting camera...")
	cam.StartListening()
	defer cam.Close()

	//TODO: init logger

	//TODO: init config

	//TODO: init app

	//TODO: Run app
}
