package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
			rbt = robot.Place(robot.Position{X: 4, Y: 4}, "")
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
