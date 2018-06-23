package main

import (
	"fmt"
	"hello/internal/app/robot"
	"log"
	"os"
	"sync"
	"time"
	"flag"
)

func RunRobot(wg *sync.WaitGroup, i int,cfg *robot.Configuration) {
	defer wg.Done()
	account := fmt.Sprintf(cfg.AccFormat, i)
	bot := robot.NewRobot(account, cfg.Password, cfg.SrvAddr)
	bot.Run()
}

func main() {
	var path string
	flag.StringVar(&path, "config", "", "please give a config path.")
	flag.StringVar(&path, "c", "F:\\star\\test\\cmd\\robot\\config.json", "please give a config path.")

	flag.Parse()

	if path == "" {
		log.Println("player give a config file use -c or -config")
		return
	}

	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file:%v", err)
		return
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("close file failed. %v", err)
		}
	}()
	log.SetOutput(f)

	cfg , err := robot.NewConfiguration(path)
	if err != nil {
		log.Printf("read configuration failed. %v",err)
		return
	}

	wg := sync.WaitGroup{}
	for i := 0; i < cfg.Num; i++ {
		time.Sleep(cfg.SleepInterval * time.Millisecond)
		wg.Add(1)
		go RunRobot(&wg, i,cfg)
	}
	wg.Wait()
}
