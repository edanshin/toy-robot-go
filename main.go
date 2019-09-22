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
	var aRobot *robot.Robot
	table := table.NewTable(5, 5)

	for {
		fmt.Print("Robot: ")

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

			// if an invalid place command is entered, get the previous valid robot
			rbt := aRobot
			aRobot = robot.Place(robot.Position{X: x, Y: y}, direction, table)

			if aRobot == nil && rbt != nil {
				aRobot = rbt
			}

			if aRobot == nil && rbt == nil {
				continue
			}
		} else if aRobot != nil {
			switch command {
			case "MOVE":
				aRobot.Move(table)
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
