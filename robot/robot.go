package robot

const (
	north = "NORTH"
	south = "SOUTH"
	east  = "EAST"
	west  = "WEST"
)

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

// Place puts a new toy robot on a table in position X,Y and facing NORTH, SOUTH, EAST or WEST
func Place(position Position, direction string) Robot {
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
	switch robot.Direction {
	case north:
		robot.Direction = west
	case south:
		robot.Direction = east
	case east:
		robot.Direction = north
	case west:
		robot.Direction = south
	}
}

// Right rotates a robot 90 degrees to the right without changing its position
func (robot *Robot) Right() {
	switch robot.Direction {
	case north:
		robot.Direction = east
	case south:
		robot.Direction = west
	case east:
		robot.Direction = south
	case west:
		robot.Direction = north
	}
}

// Report announces the position and direction of a robot
func (robot *Robot) Report() {

}
