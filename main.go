package main

import (
	"os"

	"toy-robot-go/processor"
	"toy-robot-go/robot"
)

func main() {
	r := new(robot.Robot)

	// if an argument for a file is provided, try to read and execute commands from the text file
	if len(os.Args) == 2 {
		processor.ReadText(r, os.Args[1])
	}

	// then execute commands entered via standard console input
	processor.ReadConsole(r)
}
