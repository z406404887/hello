package main

import (
	"game"
	"log"
)

func main()  {
	game,err := game.NewGame("F:\\star\\test\\cmd\\game\\config.json")
	if err != nil{
		log.Printf("create game failed. %v",err)
		return
	}
	game.Run()
}
