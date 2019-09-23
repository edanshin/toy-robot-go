// Package robot describes a toy robot
package robot

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"../table"
)

const (
	// North Direction
	North Direction = 0

	// East  Direction
	East Direction = 1

	// South Direction
	South Direction = 2

	// West  Direction
	West Direction = 3
)

// Direction defines direction the robot currently faces
type Direction int

// Robot .
type Robot struct {
	Position  Position
	Direction Direction
}

// Position of a robot .
type Position struct {
	X int
	Y int
}

// enumerate directions
func (direction Direction) String() string {
	// declare an array of directions
	directions := []string{
		"NORTH",
		"EAST",
		"SOUTH",
		"WEST",
	}

	return directions[direction]
}

// Place puts a new toy robot on a table in position X,Y and facing NORTH, SOUTH, EAST or WEST
func Place(position Position, direction Direction, table table.Table) *Robot {
	// if robot's position is not within the table, return nil
	if position.X > table.Dimensions.X-1 || position.Y > table.Dimensions.Y-1 || position.X < 0 || position.Y < 0 {
		return nil
	}

	return &Robot{
		Position:  position,
		Direction: direction,
	}
}

// Move moves a toy robot one unit forward in the direction it is currently facing
func (robot *Robot) Move(table table.Table) {
	// move the robot 1 step in the direction it currently faces,
	// only if it won't exceed table dimensions
	switch robot.Direction {
	case North:
		if robot.Position.Y < table.Dimensions.Y-1 {
			robot.Position.Y++
		}
	case South:
		if robot.Position.Y > 0 {
			robot.Position.Y--
		}
	case East:
		if robot.Position.X < table.Dimensions.X-1 {
			robot.Position.X++
		}
	case West:
		if robot.Position.X > 0 {
			robot.Position.X--
		}
	}
}

// Left rotates a robot 90 degrees to the left without changing its position
func (robot *Robot) Left() {
	if robot.Direction != North {
		robot.Direction--
	} else {
		// if robot currently faces north, set its direction to west,
		// to keep direction's number within the valid range
		robot.Direction = West
	}
}

// Right rotates a robot 90 degrees to the right without changing its position
func (robot *Robot) Right() {
	if robot.Direction != West {
		robot.Direction++
	} else {
		// if robot currently faces west, set its direction to north,
		// to keep direction's number within the valid range
		robot.Direction = North
	}
}

// Report announces current position and direction of the robot on the table
func (robot *Robot) Report() {
	fmt.Println("\n", robot.Setting()+"\n")
}

// Setting returns current position and direction of the robot on the table
func (robot *Robot) Setting() string {
	return fmt.Sprint(robot.Position.X, ",", robot.Position.Y, ",", robot.Direction.String())
}

// Display outputs the current view of the robot on the table to console
func (robot *Robot) Display(table table.Table) {
	// graphical direction
	dir := [...]string{"^", ">", "v", "<"}
	fmt.Println()

	// navigate through table dimensions
	for y := table.Dimensions.Y - 1; y >= 0; y-- {
		data := ""

		for x := 0; x < table.Dimensions.X; x++ {

			// if robot's position equals table's coordinate, show robot's direction,
			// otherwise display table's standard empty field
			if robot.Position.X == x && robot.Position.Y == y {
				data += dir[robot.Direction] + " "
			} else {
				data += "o "
			}
		}

		fmt.Println(data)
	}

	fmt.Println()
}

// Process processes commands given to robot
func Process(command string, aRobot *Robot, table table.Table) *Robot {
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

		// if an invalid place command is entered, get the previous valid robot
		rbt := aRobot
		aRobot = Place(Position{X: x, Y: y}, direction, table)

		if aRobot == nil && rbt != nil {
			aRobot = rbt
		}

		//if aRobot == nil && rbt == nil {
		//	continue
		//}
	} else if aRobot != nil {
		switch command {
		// if robot is valid, allow execution of the rest of valid commands
		case "MOVE":
			aRobot.Move(table)
		case "LEFT":
			aRobot.Left()
		case "RIGHT":
			aRobot.Right()
		case "REPORT":
			aRobot.Report()
			aRobot.Display(table)
			// fmt.Println("Invalid command entered.")
		}
	}

	return aRobot
}
