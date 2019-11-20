package main

import (
	"fmt"
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
func (robot *Robot) Place(position Position, direction Direction) bool {
	// if robot's position is not within the table, return false
	if position.X > tableWidth-1 || position.Y > tableLength-1 || position.X < 0 || position.Y < 0 {
		return false
	}

	robot.Position = position
	robot.Direction = direction
	robot.Placed = true

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
