package main

//CommandType is the kind of command
type CommandType int

//CommandType enum
const (
	DoNothing CommandType = iota
	Move
	Pickup
	Drop
	Skill
	Equip
	Attack
	Talk
)

//BaseCommand is the base class for commands
type BaseCommand struct {
	cmdType CommandType
	source  interface{}
}

//Command interface execute changes for Entities and returns the number of game ticks that occur
type Command interface {
	Execute(args ...interface{}) int
}

//Execute for the base command just returns 10 ticks
func (cmd BaseCommand) Execute(args ...interface{}) int {
	return 10
}

//MoveCommand represents a movement command
type MoveCommand struct {
	BaseCommand
}
