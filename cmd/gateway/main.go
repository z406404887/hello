package main

import (
	"flag"
	"hello/internal/app/gateway"
	"log"
	"os"
	"hello/internal/pkg/util"
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
	flag.StringVar(&path, "c", "F:\\star\\test\\cmd\\gateway\\config.json", "please give a config path.")

	flag.Parse()

	if path == "" {
		log.Println("player give a config file use -c or -config")
		return
	}

	gate, err := gateway.NewGateway(path)
	if err != nil {
		log.Printf("create gateway failed. %v", err)
		return
	}

	if err = gate.Run(); err != nil {
		log.Fatalf("run gateway failed. %v",err)
		return
	}
}
