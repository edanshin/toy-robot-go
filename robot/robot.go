package robot

// Robot .
type Robot struct {
	Position  Position
	Direction string
}

// Position of a robot .
type Position struct {
	X int
	Y int
}

// Place will put a new toy robot on a table in position X,Y and facing NORTH, SOUTH, EAST or WEST
func Place(position Position, direction string) Robot {
	robot := Robot{
		Position:  position,
		Direction: direction,
	}

	return robot
}

// Move will move a toy robot one unit forward in the direction it is currently facing
func (robot *Robot) Move() {

}

// Left will rotate a robot 90 degrees to the left without changing its position
func (robot *Robot) Left() {

}

// Right will rotate a robot 90 degrees to the right without changing its position
func (robot *Robot) Right() {

}

// Report will announce the position and direction of a robot
func (robot *Robot) Report() {

}
