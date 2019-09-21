package robot

import "fmt"

const (
	north Direction = 0
	east  Direction = 1
	south Direction = 2
	west  Direction = 3
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
func Place(position Position, direction Direction) Robot {
	robot := Robot{
		Position:  position,
		Direction: direction,
	}

	return robot
}

// Move moves a toy robot one unit forward in the direction it is currently facing
func (robot *Robot) Move() {
	switch robot.Direction {
	case north:
		robot.Position.Y++
	case south:
		robot.Position.Y--
	case east:
		robot.Position.X++
	case west:
		robot.Position.X--
	}
}

// Left rotates a robot 90 degrees to the left without changing its position
func (robot *Robot) Left() {
	if robot.Direction != north {
		robot.Direction--
	} else {
		robot.Direction = west
	}
}

// Right rotates a robot 90 degrees to the right without changing its position
func (robot *Robot) Right() {
	if robot.Direction != west {
		robot.Direction++
	} else {
		robot.Direction = north
	}
}

// Report announces the position and direction of a robot
func (robot *Robot) Report() {
	fmt.Println("Position X:", robot.Position.X, "\n", "Position Y:", robot.Position.Y, "\n", "Direction:", robot.Direction.String())
}
