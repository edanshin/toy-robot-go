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
