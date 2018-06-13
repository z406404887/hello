package main

import (
	"manager"
	"log"
)

func main() {
	mgr, err := manager.NewManager("F:\\star\\test\\cmd\\manager\\config.json")
	if err != nil {
		log.Printf("create manager failed. %v",err)
		return
	}

	mgr.Run()
}
