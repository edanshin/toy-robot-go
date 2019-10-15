// Package processor is responsible for reading and executing commands from console input
package processor

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"../robot"
)

// ReadConsole executes commands entered via standard console input
func ReadConsole(aRobot *robot.Robot) {
	fmt.Print("Robot: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		command := strings.ToUpper(scanner.Text())
		aRobot = robot.Process(command, aRobot)
	}

	ReadConsole(aRobot)
}

// ReadText reads and executes commands from a text file
func ReadText(aRobot *robot.Robot) {
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
}
