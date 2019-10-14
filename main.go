package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./processor"
	"./robot"
)

func main() {
	var aRobot *robot.Robot

	// if an argument for a file is provided, try to read and execute commands from the text file
	if len(os.Args) == 2 {
		file, err := os.Open(os.Args[1])

		if err == nil {
			defer file.Close()
			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				command := strings.ToUpper(scanner.Text())
				// output each command to terminal
				fmt.Println(command)
				aRobot = robot.Process(command, aRobot)
			}
		}
	}

	// then execute commands entered via standard console input
	processor.ReadConsole(aRobot)
}
