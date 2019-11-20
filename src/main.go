package main

import (
	"os"
)

func main() {
	robot := &Robot{}

	// if an argument for a file is provided, try to read and execute commands from the text file
	if len(os.Args) == 2 {
		ReadText(robot, os.Args[1])
	}

	// then execute commands entered via standard console input
	ReadConsole(robot)
}
