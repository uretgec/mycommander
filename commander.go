package mycommander

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// commander
type Commander struct {
	currentPath string
	timeout     int64 // minutes

	debugMode bool

	commands map[string]*Command
}

// create new commander
func NewCommander(commands ...*Command) *Commander {
	c := &Commander{
		currentPath: "",
		commands:    map[string]*Command{},
		timeout:     0, // minutes
		debugMode:   false,
	}

	// update commander current working path
	currentPath, _ := c.findCurrentPath()
	c.currentPath = currentPath

	if len(commands) > 0 {
		for _, cmd := range commands {
			c.commands[cmd.name] = cmd
		}
	}

	return c
}

// activate debug mode
func (c *Commander) isDebugModeActive() bool {
	return c.debugMode
}

// update timeout: minutes
func (c *Commander) UpdateTimeout(timeout int64) {
	c.timeout = int64(time.Duration(timeout) * time.Minute / time.Millisecond)
}

// reset timeout
func (c *Commander) ResetTimeout() {
	c.timeout = 0
}

// add command/s
func (c *Commander) add(commands ...*Command) {
	for _, cmd := range commands {
		c.commands[cmd.name] = cmd
	}
}

// remove command/s
func (c *Commander) remove(names ...string) {
	for _, name := range names {
		delete(c.commands, name)
	}
}

// valid command
func (c *Commander) valid(cmd string) bool {
	_, ok := c.commands[cmd]
	return ok
}

// find current working path - look like pwd
func (c *Commander) findCurrentPath() (string, error) {
	return os.Getwd()
}

// merge multiple commands
/*func (c *Commander) mergeMultipleCommands(commands ...string) string {
	return strings.Join(commands, "; ")
}*/

// run forest run
func (c *Commander) run(name string, args ...string) ([]byte, error) {
	if !c.valid(name) {
		return nil, errors.New("unknown command")
	}

	cmd := c.commands[name]
	if len(cmd.args) != len(args) {
		return nil, errors.New("not enough arguments")
	}

	var sh *exec.Cmd

	if c.timeout > 0 {
		// timeot option
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.timeout)*time.Second)
		defer cancel()

		sh = exec.CommandContext(ctx, "/bin/bash", "-c", cmd.generate(args...))
	} else {
		sh = exec.Command("/bin/bash", "-c", cmd.generate(args...))
	}

	sh.Dir = c.currentPath
	if c.isDebugModeActive() {
		fmt.Println("Commands: ", sh.String(), c.timeout, c.currentPath)
	}

	out, err := sh.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return out, nil
}
