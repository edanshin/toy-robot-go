package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ReadConsole executes commands entered via standard console input
func ReadConsole(robot *Robot) {
	fmt.Print("Robot: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		command := strings.ToUpper(scanner.Text())
		robot = Process(command, robot)
	}

	ReadConsole(robot)
}

// ReadText reads and executes commands from a text file
func ReadText(robot *Robot, filepath string) bool {
	// if an argument for a file is provided, try to read and execute commands from the text file
	file, err := os.Open(filepath)

	if err != nil {
		return false
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		command := strings.ToUpper(scanner.Text())
		// output each command to terminal
		fmt.Println(command)
		robot = Process(command, robot)
	}

	return true
}

// Process processes commands given to robot
func Process(command string, robot *Robot) *Robot {
	if strings.HasPrefix(command, "PLACE") {
		// check if submitted command is a valid PLACE command
		regex := regexp.MustCompile(`PLACE \d,\d,(NORTH|SOUTH|EAST|WEST)`).MatchString(command)
		if regex {
			// get command's parameters
			cmd := strings.Split(command, " ")
			cmd = strings.Split(cmd[1], ",")

			// get command's coordinate
			x, _ := strconv.Atoi(cmd[0])
			y, _ := strconv.Atoi(cmd[1])

			// extract direction
			var direction Direction

			switch cmd[2] {
			case "NORTH":
				direction = North
			case "EAST":
				direction = East
			case "SOUTH":
				direction = South
			case "WEST":
				direction = West
			}

			robot.Place(Position{X: x, Y: y}, direction)
		} else {
			fmt.Println("Invalid PLACE command entered.")
		}
	} else if command == "MOVE" && robot.Placed {
		robot.Move()
	} else if command == "LEFT" && robot.Placed {
		robot.Left()
	} else if command == "RIGHT" && robot.Placed {
		robot.Right()
	} else if command == "REPORT" && robot.Placed {
		robot.Report()
		robot.Display()
	} else if command == "EXIT" {
		os.Exit(0)
	} else if !robot.Placed && command != "EXIT" {
		fmt.Println("Error: robot is not placed.")
	} else {
		fmt.Println("Invalid command entered.")
	}

	return robot
}
