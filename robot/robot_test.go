package robot

import (
	"testing"

	"../table"
)

// TestPlace tests robot's PLACE command
func TestPlace(t *testing.T) {
	var aRobot *Robot
	table := table.NewTable(5, 5)
	command := "PLACE 2,3,EAST"

	aRobot = Process(command, aRobot, table)

	if aRobot == nil {
		t.Error("PLACE command test failed.")
	}

	if aRobot.Setting() != "2,3,EAST" {
		t.Error("PLACE command test failed, expected output 2,3,EAST, got ", aRobot.Setting())
	}
}

// TestMove tests robot's MOVE command
func TestMove(t *testing.T) {
	var aRobot *Robot
	table := table.NewTable(5, 5)
	commands := []string{"PLACE 2,3,EAST", "MOVE"}

	for _, command := range commands {
		aRobot = Process(command, aRobot, table)
	}

	if aRobot == nil {
		t.Error("MOVE command test failed.")
	}

	if aRobot.Setting() != "3,3,EAST" {
		t.Error("MOVE command test failed, expected output 3,3,EAST, got ", aRobot.Setting())
	}
}

// TestLeft tests robot's LEFT command
func TestLeft(t *testing.T) {
	var aRobot *Robot
	table := table.NewTable(5, 5)
	commands := []string{"PLACE 2,3,EAST", "LEFT"}

	for _, command := range commands {
		aRobot = Process(command, aRobot, table)
	}

	if aRobot == nil {
		t.Error("LEFT command test failed.")
	}

	if aRobot.Setting() != "2,3,NORTH" {
		t.Error("LEFT command test failed, expected output 2,3,NORTH, got ", aRobot.Setting())
	}
}

// TestRight tests robot's RIGHT command
func TestRight(t *testing.T) {
	var aRobot *Robot
	table := table.NewTable(5, 5)
	commands := []string{"PLACE 2,3,EAST", "RIGHT"}

	for _, command := range commands {
		aRobot = Process(command, aRobot, table)
	}

	if aRobot == nil {
		t.Error("RIGHT command test failed.")
	}

	if aRobot.Setting() != "2,3,SOUTH" {
		t.Error("RIGHT command test failed, expected output 2,3,SOUTH, got ", aRobot.Setting())
	}
}

// TestCaseA tests Example a from problem description
func TestCaseA(t *testing.T) {
	var aRobot *Robot
	table := table.NewTable(5, 5)
	commands := []string{"PLACE 0,0,NORTH", "MOVE", "REPORT"}

	for _, command := range commands {
		aRobot = Process(command, aRobot, table)
	}

	if aRobot == nil {
		t.Error("Example a test failed.")
	}

	if aRobot.Setting() != "0,1,NORTH" {
		t.Error("Example a test failed, expected output 0,1,NORTH, got ", aRobot.Setting())
	}
}

// TestCaseB tests Example b from problem description
func TestCaseB(t *testing.T) {
	var aRobot *Robot
	table := table.NewTable(5, 5)
	commands := []string{"PLACE 0,0,NORTH", "LEFT", "REPORT"}

	for _, command := range commands {
		aRobot = Process(command, aRobot, table)
	}

	if aRobot == nil {
		t.Error("Example b test failed.")
	}

	if aRobot.Setting() != "0,0,WEST" {
		t.Error("Example b test failed, expected output 0,0,WEST, got ", aRobot.Setting())
	}
}

// TestCaseC tests Example c from problem description
func TestCaseC(t *testing.T) {
	var aRobot *Robot
	table := table.NewTable(5, 5)
	commands := []string{"PLACE 1,2,EAST", "MOVE", "MOVE", "LEFT", "MOVE", "REPORT"}

	for _, command := range commands {
		aRobot = Process(command, aRobot, table)
	}

	if aRobot == nil {
		t.Error("Example c test failed.")
	}

	if aRobot.Setting() != "3,3,NORTH" {
		t.Error("Example c test failed, expected output 3,3,NORTH, got ", aRobot.Setting())
	}
}
