package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./robot"
	"./table"
)

func main() {
	var aRobot *robot.Robot
	table := table.NewTable(5, 5)

	// if an argument for a file is provided, try to read and execute commands from the file
	if len(os.Args) == 2 {
		file, err := os.Open(os.Args[1])

		if err != nil {
			return
		}

		defer file.Close()
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			command := strings.ToUpper(scanner.Text())
			aRobot = robot.Process(command, aRobot, table)
		}
	} else if len(os.Args) == 1 {
		// if no arguments are provided, execute commands entered via standard console input
		for {
			fmt.Print("Robot: ")

			scanner := bufio.NewScanner(os.Stdin)

			if !scanner.Scan() {
				continue
			}

			command := strings.ToUpper(scanner.Text())
			aRobot = robot.Process(command, aRobot, table)
		}
	}
}
