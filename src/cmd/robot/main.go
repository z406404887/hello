package main

import (
	"robot"
)

func main() {
	bot := robot.NewRobot("111", "222", "ws://127.0.0.1:11000")
	bot.Run()
}
