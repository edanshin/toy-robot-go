# Toy Robot Simulator
=====================
### Description
---------------
- The command line application is a simulation of a toy robot moving on a square tabletop,
  of dimensions 5 units x 5 units.
- There are no other obstructions on the table surface.
- The robot is free to roam around the surface of the table.

Any movement that would result in the
robot falling from the table is prevented, however further valid
movement commands are still allowed.

### Commands
The application can read commands of the following (textual) form:

- PLACE X,Y,F - puts the toy robot on the table in position X,Y and facing NORTH,
  SOUTH, EAST or WEST.

- MOVE - moves the toy robot one unit forward in the direction it is
  currently facing.

- LEFT - rotates the robot 90 degrees to the left
  without changing the position of the robot.

- RIGHT - rotates the robot 90 degrees to the right
  without changing the position of the robot.

- REPORT - announces the X,Y and F (direction) of the robot.

### Conditions
- The first valid command to the robot is a valid PLACE command, after that, any
  sequence of commands can be issued, in any order, including another PLACE
  command. The application discards all commands in the sequence until
  a valid PLACE command has been executed.

- The origin (0,0) is the SOUTH WEST most corner.

- A robot that is not on the table ignores the MOVE, LEFT, RIGHT
  and REPORT commands.

### Test data
- Test data to exercise the application is provided in the form of unit tests
(robot/robot_test.go).

### Input
Input is possible via standard console input as well as via file input.
Commands will be converted to uppercase automatically and are not case sensitive.
If text file with commands is not found or is empty, input mode will switch
to standard console input.

### Output
The program has the ability to output graphical representation of a tabletop
with a robot on it to the console. This also includes robot's direction.

### Design
Object-oriented approach has been used for the development of this application.

- Package "robot" contains all the functions necessary for robot's functionality
(Place, Move, Left, Right, Report).
- Package "table" describes a tabletop, where a robot is moving.
- The file "main.go" is responsible for program execution.

### Development tools
   1) Go 1.13
   2) VS Code

### Build and run the program
- Pre-compiled version of the program is provided.
- To build the app, execute "go build" command in the app folder.
- To run the app, execute "./toy-robot" command in the app folder.
- To run the app with predefined commands, execute "./toy-robot [commands file path]"
command in the app folder. The commands file should be a text file with a list of commands,
delimited by new line.
