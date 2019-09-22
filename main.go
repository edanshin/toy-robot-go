package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"./robot"
	"./table"
)

func main() {
	var aRobot robot.Robot
	robotCreated := false

	for {
		fmt.Print("Command: ")

		scanner := bufio.NewScanner(os.Stdin)

		if !scanner.Scan() {
			continue
		}

		command := strings.ToUpper(scanner.Text())

		regex := regexp.MustCompile(`PLACE \d,\d,(NORTH|SOUTH|EAST|WEST)`).MatchString(command)
		if regex {
			cmd := strings.Split(command, " ")
			cmd = strings.Split(cmd[1], ",")

			x, _ := strconv.Atoi(cmd[0])
			y, _ := strconv.Atoi(cmd[1])

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

			table := table.NewTable(5, 5)

			// make sure our new robot's position is within the table surface
			if x > table.Dimensions.X || y > table.Dimensions.Y || x < 0 || y < 0 {
				continue
			}

			aRobot = robot.Place(robot.Position{X: x, Y: y}, direction)

			if !robotCreated {
				robotCreated = true
			}

			aRobot.Report()
		} else if robotCreated {
			switch command {
			case "MOVE":
				aRobot.Move()
			case "LEFT":
				aRobot.Left()
			case "RIGHT":
				aRobot.Right()
			case "REPORT":
				aRobot.Report()

				// fmt.Println("Invalid command entered.")
			}
		}
	}
}
