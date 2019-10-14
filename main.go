package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	for {
		fmt.Print("Robot: ")
		scanner := bufio.NewScanner(os.Stdin)

		if !scanner.Scan() {
			continue
		}

		command := strings.ToUpper(scanner.Text())
		aRobot = robot.Process(command, aRobot)
	}
}
