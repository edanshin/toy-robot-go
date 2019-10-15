// Package robot describes a toy robot
package robot

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

const (
	tableWidth  = 5
	tableLength = 5
)

// Direction defines direction the robot currently faces
type Direction int

// Robot .
type Robot struct {
	Position  Position
	Direction Direction
	Placed    bool
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
func Place(position Position, direction Direction, r *Robot) bool {
	// if robot's position is not within the table, return false
	if position.X > tableWidth-1 || position.Y > tableLength-1 || position.X < 0 || position.Y < 0 {
		return false
	}

	r.Position = position
	r.Direction = direction
	r.Placed = true

	return true
}

// Move moves a toy robot one unit forward in the direction it is currently facing
func (robot *Robot) Move() {
	// move the robot 1 step in the direction it currently faces,
	// only if it won't exceed table dimensions
	switch robot.Direction {
	case North:
		if robot.Position.Y < tableLength-1 {
			robot.Position.Y++
		}
	case South:
		if robot.Position.Y > 0 {
			robot.Position.Y--
		}
	case East:
		if robot.Position.X < tableWidth-1 {
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
	fmt.Println("\n" + robot.Setting())
}

// Setting returns current position and direction of the robot on the table
func (robot *Robot) Setting() string {
	return fmt.Sprint(robot.Position.X, ",", robot.Position.Y, ",", robot.Direction.String())
}

// Display outputs the current view of the robot on the table to console
func (robot *Robot) Display() {
	// graphical direction
	dir := [...]string{"^", ">", "v", "<"}
	fmt.Println()

	// navigate through table dimensions
	for y := tableLength - 1; y >= 0; y-- {
		data := ""

		for x := 0; x < tableWidth; x++ {

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
func Process(command string, aRobot *Robot) *Robot {
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

			Place(Position{X: x, Y: y}, direction, aRobot)
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
	} else {
		fmt.Println("Invalid command entered.")
	}

	return aRobot
}
