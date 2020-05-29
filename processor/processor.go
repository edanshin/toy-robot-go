package processor

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"toy-robot-go/robot"
)

// ReadConsole executes commands entered via standard console input
func ReadConsole(r *robot.Robot) {
	fmt.Print("Robot: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		command := strings.ToLower(scanner.Text())
		Process(command, r)
	}

	ReadConsole(r)
}

// ReadText reads and executes commands from a text file
func ReadText(r *robot.Robot, filepath string) bool {
	// if an argument for a file is provided, try to read and execute commands from the text file
	file, err := os.Open(filepath)

	if err != nil {
		return false
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		command := strings.ToLower(scanner.Text())
		// output each command to terminal
		fmt.Println(command)
		Process(command, r)
	}

	return true
}

// Process processes commands given to robot
func Process(command string, r *robot.Robot) {
	switch {
	case strings.HasPrefix(command, "place"):
		// check if submitted command is a valid PLACE command
		regex := regexp.MustCompile(`place \d,\d,(north|south|east|west)`).MatchString(command)
		if regex {
			// get command's parameters
			cmd := strings.Split(command, " ")
			cmd = strings.Split(cmd[1], ",")

			// get command's coordinate
			x, _ := strconv.Atoi(cmd[0])
			y, _ := strconv.Atoi(cmd[1])

			// extract direction
			var direction robot.Direction

			switch cmd[2] {
			case "north":
				direction = robot.North
			case "east":
				direction = robot.East
			case "south":
				direction = robot.South
			case "west":
				direction = robot.West
			}

			r.Place(robot.Position{X: x, Y: y}, direction)
		} else {
			fmt.Println("Invalid PLACE command entered.")
		}

		break

	case command == "move" && r.Placed:
		r.Move()
		break

	case command == "left" && r.Placed:
		r.Left()
		break

	case command == "right" && r.Placed:
		r.Right()
		break

	case command == "report" && r.Placed:
		r.Report()
		r.Display()
		break

	case command == "exit":
		os.Exit(0)
		break

	case !r.Placed && command != "exit":
		fmt.Println("Error: robot is not placed.")
		break

	default:
		fmt.Println("Invalid command entered.")
		break
	}
}
