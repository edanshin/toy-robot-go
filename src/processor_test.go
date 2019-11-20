package main

import (
	"testing"
)

// TestPlace tests robot's PLACE command
func TestPlace(t *testing.T) {
	robot := &Robot{}
	command := "PLACE 2,3,EAST"

	robot = Process(command, robot)

	if !robot.Placed {
		t.Error("PLACE command test failed.")
	}

	if robot.Setting() != "2,3,EAST" {
		t.Error("PLACE command test failed, expected output 2,3,EAST, got ", robot.Setting())
	} else {
		t.Log("PLACE command test successful, got ", robot.Setting())
	}
}

// TestMove tests robot's MOVE command
func TestMove(t *testing.T) {
	robot := &Robot{}
	commands := []string{"PLACE 2,3,EAST", "MOVE"}

	for _, command := range commands {
		robot = Process(command, robot)
	}

	if !robot.Placed {
		t.Error("MOVE command test failed.")
	}

	if robot.Setting() != "3,3,EAST" {
		t.Error("MOVE command test failed, expected output 3,3,EAST, got ", robot.Setting())
	} else {
		t.Log("MOVE command test successful, got ", robot.Setting())
	}
}

// TestLeft tests robot's LEFT command
func TestLeft(t *testing.T) {
	robot := &Robot{}
	commands := []string{"PLACE 2,3,EAST", "LEFT"}

	for _, command := range commands {
		robot = Process(command, robot)
	}

	if !robot.Placed {
		t.Error("LEFT command test failed.")
	}

	if robot.Setting() != "2,3,NORTH" {
		t.Error("LEFT command test failed, expected output 2,3,NORTH, got ", robot.Setting())
	} else {
		t.Log("LEFT command test successful, got ", robot.Setting())
	}
}

// TestRight tests robot's RIGHT command
func TestRight(t *testing.T) {
	robot := &Robot{}
	commands := []string{"PLACE 2,3,EAST", "RIGHT"}

	for _, command := range commands {
		robot = Process(command, robot)
	}

	if !robot.Placed {
		t.Error("RIGHT command test failed.")
	}

	if robot.Setting() != "2,3,SOUTH" {
		t.Error("RIGHT command test failed, expected output 2,3,SOUTH, got ", robot.Setting())
	} else {
		t.Log("RIGHT command test successful, got ", robot.Setting())
	}
}

// TestReadText tests execution of commands from a text file
func TestReadText(t *testing.T) {
	robot := &Robot{}

	ReadText(robot, "./commands.txt")

	if !robot.Placed {
		t.Error("ReadText test failed.")
	}

	if robot.Setting() != "3,3,NORTH" {
		t.Error("ReadText test failed, expected output 3,3,NORTH, got ", robot.Setting())
	} else {
		t.Log("ReadText test successful, got ", robot.Setting())
	}
}

// TestCaseA tests Example a from problem description
func TestCaseA(t *testing.T) {
	robot := &Robot{}
	commands := []string{"PLACE 0,0,NORTH", "MOVE", "REPORT"}

	for _, command := range commands {
		robot = Process(command, robot)
	}

	if !robot.Placed {
		t.Error("Example a test failed.")
	}

	if robot.Setting() != "0,1,NORTH" {
		t.Error("Example a test failed, expected output 0,1,NORTH, got ", robot.Setting())
	} else {
		t.Log("Example a test successful, got ", robot.Setting())
	}
}

// TestCaseB tests Example b from problem description
func TestCaseB(t *testing.T) {
	robot := &Robot{}
	commands := []string{"PLACE 0,0,NORTH", "LEFT", "REPORT"}

	for _, command := range commands {
		robot = Process(command, robot)
	}

	if !robot.Placed {
		t.Error("Example b test failed.")
	}

	if robot.Setting() != "0,0,WEST" {
		t.Error("Example b test failed, expected output 0,0,WEST, got ", robot.Setting())
	} else {
		t.Log("Example b test successful, got ", robot.Setting())
	}
}

// TestCaseC tests Example c from problem description
func TestCaseC(t *testing.T) {
	robot := &Robot{}
	commands := []string{"PLACE 1,2,EAST", "MOVE", "MOVE", "LEFT", "MOVE", "REPORT"}

	for _, command := range commands {
		robot = Process(command, robot)
	}

	if !robot.Placed {
		t.Error("Example c test failed.")
	}

	if robot.Setting() != "3,3,NORTH" {
		t.Error("Example c test failed, expected output 3,3,NORTH, got ", robot.Setting())
	} else {
		t.Log("Example c test successful, got ", robot.Setting())
	}
}
