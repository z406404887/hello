package main

import (
	"log"
	"login"
)

func main() {
	login,err := login.NewLogin("F:\\star\\test\\cmd\\login\\config.json")
	if err != nil{
		log.Printf("create login failed. %v",err)
		return
	}
	login.Run()
}
