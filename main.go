package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"./robot"
)

func main() {
	var aRobot robot.Robot

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print("Command: ")

		command := strings.ToUpper(scanner.Text())

		regex := regexp.MustCompile(`PLACE \d,\d,(NORTH|SOUTH|EAST|WEST)`).MatchString(command)
		if regex {
			cmd := strings.Split(command, " ")
			cmd = strings.Split(cmd[1], ",")

			x, _ := strconv.Atoi(cmd[0])
			y, _ := strconv.Atoi(cmd[1])
			direction := cmd[2]

			aRobot = robot.Place(robot.Position{X: x, Y: y}, direction)
		}

		switch command {
		case "MOVE":
			aRobot.Move()
		case "LEFT":
			aRobot.Left()
		case "RIGHT":
			aRobot.Right()
		case "REPORT":
			aRobot.Report()
		default:
			fmt.Println("Invalid command entered.")
		}
	}
}
