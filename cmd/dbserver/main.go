package main

import (
	"flag"
	"log"
	"os"
	"starter-kit/internal/app/db"
	"starter-kit/internal/pkg/util"
)

func main() {
	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file:%v", err)
		return
	}

	defer util.Close(f)
	log.SetOutput(f)

	var path string
	flag.StringVar(&path, "config", "", "please give a config path.")
	flag.StringVar(&path, "c", "F:\\star\\test\\cmd\\dbserver\\config.json", "please give a config path.")

	flag.Parse()
	if path == "" {
		log.Println("player give a config file use -c or -config")
		return
	}

	srv, err := db.NewDbServer(path)
	if err != nil {
		log.Printf("create db server failed. %v", err)
		return
	}
	if err = srv.Run(); err != nil {
		log.Fatalf("fatal error, exiting.%v", err)
	}
}
