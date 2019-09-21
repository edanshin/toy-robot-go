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
	if width < 1 {
		width = 4
	}

	if length < 1 {
		length = 4
	}

	table := Table{
		Dimensions: Dimensions{
			X: width,
			Y: length,
		},
	}

	return table
}
