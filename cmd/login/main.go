package main

import (
	"flag"
	"hello/internal/app/login"
	"log"
)

func main() {
	var path string
	flag.StringVar(&path, "config", "", "please give a config path.")
	flag.StringVar(&path, "c", "F:\\star\\test\\cmd\\login\\config.json", "please give a config path.")

	flag.Parse()

	if path == "" {
		log.Println("player give a config file use -c or -config")
		return
	}

	login, err := login.NewLogin(path)
	if err != nil {
		log.Printf("create login failed. %v", err)
		return
	}
	if err = login.Run(); err != nil {
		log.Fatalf("run login failed. %v", err)
		return
	}
}
