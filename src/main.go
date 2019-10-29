package main

import (
	"os"

	"./processor"
	"./robot"
)

func main() {
	aRobot := &robot.Robot{}

	// if an argument for a file is provided, try to read and execute commands from the text file
	if len(os.Args) == 2 {
		processor.ReadText(aRobot, os.Args[1])
	}

	// then execute commands entered via standard console input
	processor.ReadConsole(aRobot)
}
