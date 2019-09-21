package table

// Table .
type Table struct {
	Dimensions Dimensions
}

// Dimensions of a table
type Dimensions struct {
	X int
	Y int
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
		Dimensions: Dimensions{
			X: width,
			Y: length,
		},
	}

	return table
}
