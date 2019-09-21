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
	var rbt robot.Robot

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

			rbt = robot.Place(robot.Position{X: x, Y: y}, direction)
		}

		switch command {
		case "MOVE":
			rbt.Move()
		case "LEFT":
			rbt.Left()
		case "RIGHT":
			rbt.Right()
		case "REPORT":
			rbt.Report()
		default:
			fmt.Println("Invalid command entered.")
		}
	}
}
