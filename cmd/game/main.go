package main

import (
	"flag"
	"hello/internal/app/game"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file:%v", err)
		return
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("close file failed. %v",err)
		}
	}()
	log.SetOutput(f)

	var path string
	flag.StringVar(&path, "config", "", "please give a config path.")
	flag.StringVar(&path, "c", "F:\\star\\test\\cmd\\game\\config.json", "please give a config path.")

	flag.Parse()

	if path == "" {
		log.Println("player give a config file use -c or -config")
		return
	}
	game, err := game.NewGame(path)
	if err != nil {
		log.Printf("create game failed. %v", err)
		return
	}
	game.Run()
}
