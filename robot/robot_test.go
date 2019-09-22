package robot

import (
	"testing"

	"../table"
)

func TestPlace(t *testing.T) {
	var aRobot *Robot
	table := table.NewTable(5, 5)
	command := "PLACE 2,3,EAST"

	aRobot = Process(command, aRobot, table)

	if aRobot == nil {
		t.Error("")
	}

	if aRobot.Setting() != "2,3,EAST" {
		t.Error("")
	}
}

func TestMove(t *testing.T) {
	var aRobot *Robot
	table := table.NewTable(5, 5)
	commands := []string{"PLACE 2,3,EAST", "MOVE"}

	for _, command := range commands {
		aRobot = Process(command, aRobot, table)
	}

	if aRobot == nil {
		t.Error("")
	}

	if aRobot.Setting() != "3,3,EAST" {
		t.Error("")
	}
}

func TestLeft(t *testing.T) {
	var aRobot *Robot
	table := table.NewTable(5, 5)
	commands := []string{"PLACE 2,3,EAST", "LEFT"}

	for _, command := range commands {
		aRobot = Process(command, aRobot, table)
	}

	if aRobot == nil {
		t.Error("")
	}

	if aRobot.Setting() != "2,3,NORTH" {
		t.Error("")
	}
}

func TestRight(t *testing.T) {
	var aRobot *Robot
	table := table.NewTable(5, 5)
	commands := []string{"PLACE 2,3,EAST", "RIGHT"}

	for _, command := range commands {
		aRobot = Process(command, aRobot, table)
	}

	if aRobot == nil {
		t.Error("")
	}

	if aRobot.Setting() != "2,3,SOUTH" {
		t.Error("")
	}
}

func TestCaseA(t *testing.T) {

}

func TestCaseB(t *testing.T) {

}

func TestCaseC(t *testing.T) {

}
