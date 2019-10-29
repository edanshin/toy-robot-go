package main

import (
	"./processor"
	"./robot"
)

func main() {
	aRobot := &robot.Robot{}

	// if an argument for a file is provided, try to read and execute commands from the text file
	processor.ReadText(aRobot)

	// then execute commands entered via standard console input
	processor.ReadConsole(aRobot)
}
