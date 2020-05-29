package robot

import (
	"fmt"
)

// Direction defines direction the robot currently faces
type Direction uint

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
func (r *Robot) Place(position Position, direction Direction) bool {
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
func (r *Robot) Move() {
	// move the robot 1 step in the direction it currently faces,
	// only if it won't exceed table dimensions
	switch r.Direction {
	case North:
		if r.Position.Y < tableLength-1 {
			r.Position.Y++
		}
	case South:
		if r.Position.Y > 0 {
			r.Position.Y--
		}
	case East:
		if r.Position.X < tableWidth-1 {
			r.Position.X++
		}
	case West:
		if r.Position.X > 0 {
			r.Position.X--
		}
	}
}

// Left rotates a robot 90 degrees to the left without changing its position
func (r *Robot) Left() {
	if r.Direction != North {
		r.Direction--
	} else {
		// if robot currently faces north, set its direction to west,
		// to keep direction's number within the valid range
		r.Direction = West
	}
}

// Right rotates a robot 90 degrees to the right without changing its position
func (r *Robot) Right() {
	if r.Direction != West {
		r.Direction++
	} else {
		// if robot currently faces west, set its direction to north,
		// to keep direction's number within the valid range
		r.Direction = North
	}
}

// Report announces current position and direction of the robot on the table
func (r *Robot) Report() {
	fmt.Print("\n", r.Setting())
}

// Setting returns current position and direction of the robot on the table
func (r *Robot) Setting() string {
	return fmt.Sprint(r.Position.X, ",", r.Position.Y, ",", r.Direction.String())
}

// Display outputs the current view of the robot on the table to console
func (r *Robot) Display() {
	// graphical direction
	dir := [...]string{"^", ">", "v", "<"}
	fmt.Println()

	// navigate through table dimensions
	for y := tableLength - 1; y >= 0; y-- {
		data := ""

		for x := 0; x < tableWidth; x++ {
			// if robot's position equals table's coordinate, show robot's direction,
			// otherwise display table's standard empty field
			if r.Position.X == x && r.Position.Y == y {
				data += dir[r.Direction] + " "
			} else {
				data += "o "
			}
		}

		fmt.Println(data)
	}

	fmt.Println()
}
