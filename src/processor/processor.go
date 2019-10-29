// Package processor is responsible for reading and executing commands from console input
package processor

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"../robot"
)

// ReadConsole executes commands entered via standard console input
func ReadConsole(aRobot *robot.Robot) {
	fmt.Print("Robot: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		command := strings.ToUpper(scanner.Text())
		aRobot = Process(command, aRobot)
	}

	ReadConsole(aRobot)
}

// ReadText reads and executes commands from a text file
func ReadText(aRobot *robot.Robot, filepath string) bool {
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
		aRobot = Process(command, aRobot)
	}

	return true
}

// Process processes commands given to robot
func Process(command string, aRobot *robot.Robot) *robot.Robot {
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
			var direction robot.Direction

			switch cmd[2] {
			case "NORTH":
				direction = robot.North
			case "EAST":
				direction = robot.East
			case "SOUTH":
				direction = robot.South
			case "WEST":
				direction = robot.West
			}

			aRobot.Place(robot.Position{X: x, Y: y}, direction)
		} else {
			fmt.Println("Invalid PLACE command entered.")
		}
	} else if command == "MOVE" && aRobot.Placed {
		aRobot.Move()
	} else if command == "LEFT" && aRobot.Placed {
		aRobot.Left()
	} else if command == "RIGHT" && aRobot.Placed {
		aRobot.Right()
	} else if command == "REPORT" && aRobot.Placed {
		aRobot.Report()
		aRobot.Display()
	} else if command == "EXIT" {
		os.Exit(0)
	} else if !aRobot.Placed && command != "EXIT" {
		fmt.Println("Error: robot is not placed.")
	} else {
		fmt.Println("Invalid command entered.")
	}

	return aRobot
}
