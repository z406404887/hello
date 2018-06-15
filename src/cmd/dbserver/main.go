package main

import (
	"db"
	"log"
)

func main(){
	srv, err := db.NewDbServer("F:\\star\\test\\cmd\\dbserver\\config.json")
	if err != nil{
		log.Printf("create db server failed. %v",err)
		return
	}
	srv.Run()
}
