package table

// Table .
type Table struct {
	dimensions Dimensions
}

// Dimensions of table
type Dimensions struct {
	x int
	y int
}

// NewTable creates a new table
func NewTable(width, length int) Table {
	if width < 2 {
		width = 5
	}

	if length < 2 {
		length = 5
	}

	table := Table{
		dimensions: Dimensions{
			x: width,
			y: length,
		},
	}

	return table
}
