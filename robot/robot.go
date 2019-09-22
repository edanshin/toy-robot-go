package robot

import (
	"fmt"

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

// Direction defines current direction the robot currently faces
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
	if position.X > table.Dimensions.X || position.Y > table.Dimensions.Y || position.X < 0 || position.Y < 0 {
		return nil
	}

	return &Robot{
		Position:  position,
		Direction: direction,
	}
}

// Move moves a toy robot one unit forward in the direction it is currently facing
func (robot *Robot) Move(table table.Table) {
	switch robot.Direction {
	case North:
		if robot.Position.Y < table.Dimensions.Y {
			robot.Position.Y++
		}
	case South:
		if robot.Position.Y > 0 {
			robot.Position.Y--
		}
	case East:
		if robot.Position.X < table.Dimensions.X {
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
		robot.Direction = West
	}
}

// Right rotates a robot 90 degrees to the right without changing its position
func (robot *Robot) Right() {
	if robot.Direction != West {
		robot.Direction++
	} else {
		robot.Direction = North
	}
}

// Report announces the position and direction of a robot
func (robot *Robot) Report() {
	fmt.Println("\n", "Position X:", robot.Position.X, "\n", "Position Y:", robot.Position.Y, "\n", "Direction:", robot.Direction.String()+"\n")
}
