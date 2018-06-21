package main

import (
	"fmt"
	"hello/internal/app/robot"
	"log"
	"os"
	"sync"
	"time"
)

func RunRobot(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	id := fmt.Sprintf("%s-%05d", "robot", i)
	bot := robot.NewRobot(id, "123456", "ws://47.98.100.204:11000")
	bot.Run()
}

func main() {
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

	wg := sync.WaitGroup{}
	for i := 0; i < 500; i++ {
		time.Sleep(500 * time.Millisecond)
		wg.Add(1)
		go RunRobot(&wg, i)
	}
	wg.Wait()
}
