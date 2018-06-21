package main

import (
	"flag"
	"hello/internal/app/manager"
	"log"
)

func main() {
	var path string
	flag.StringVar(&path, "config", "", "please give a config path.")
	flag.StringVar(&path, "c", "F:\\star\\test\\cmd\\manager\\config.json", "please give a config path.")

	flag.Parse()

	if path == "" {
		log.Println("player give a config file use -c or -config")
		return
	}
	mgr, err := manager.NewManager(path)
	if err != nil {
		log.Printf("create manager failed. %v", err)
		return
	}

	if err = mgr.Run(); err != nil {
		log.Fatalf("run manager failed.%v", err)
	}
}
