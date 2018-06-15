package main

import (
	"db"
	"log"
	"flag"
)

func main(){
	var path string
	flag.StringVar(&path, "config", "", "please give a config path.")
	flag.StringVar(&path, "c", "F:\\star\\test\\cmd\\dbserver\\config.json", "please give a config path.")

	flag.Parse()
	if path == "" {
		log.Println("player give a config file use -c or -config")
		return
	}

	srv, err := db.NewDbServer(path)
	if err != nil{
		log.Printf("create db server failed. %v",err)
		return
	}
	srv.Run()
}
