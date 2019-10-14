package processor

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"../robot"
)

// ReadConsole executes commands entered via standard console input
func ReadConsole(aRobot *robot.Robot) {
	fmt.Print("Robot: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		command := strings.ToUpper(scanner.Text())
		aRobot = robot.Process(command, aRobot)
	}

	ReadConsole(aRobot)
}
