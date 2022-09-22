package mycommander

import (
	"fmt"
	"strings"
)

// delimeter
type Delimeter struct {
	start string
	end   string
}

// command
type Command struct {
	name      string
	template  string
	delimeter Delimeter
	args      []string
}

// create new command
func NewCommand(name, template string, args ...string) *Command {
	return &Command{
		name:      name,
		template:  template,
		delimeter: Delimeter{start: "[", end: "]"},
		args:      args,
	}
}

// update delimeter
func (c *Command) UpdateDelimeter(start, end string) {
	c.delimeter = Delimeter{
		start: start,
		end:   end,
	}
}

// generate command
func (c *Command) generate(args ...string) string {
	if len(c.args) > 0 {
		replacerMap := []string{}
		for in, arg := range c.args {
			replacerMap = append(replacerMap, fmt.Sprintf("%s%s%s", c.delimeter.start, arg, c.delimeter.end), args[in])
		}

		replacer := strings.NewReplacer(replacerMap...)
		return replacer.Replace(c.template)
	}

	return c.template
}
