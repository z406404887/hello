package main

import (
	"gateway"
	"log"
)

func main()  {
	gate,err := gateway.NewGateway("F:\\star\\test\\cmd\\gateway\\config.json")
	if err != nil{
		log.Printf("create gateway failed. %v",err)
		return
	}
	gate.Run()
}

