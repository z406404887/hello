package main

import (
	"robot"
)

func main() {
	bot := robot.NewRobot("111","222","ws://192.168.2.52:11000")
	bot.Run()
}
