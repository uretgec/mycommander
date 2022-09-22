package mycommander

// my command manager
type MyCommander struct {
	transports *Transporter
	hosts      *Hoster
	commands   *Commander
}

// create new mycommander
func NewMyCommander(transports *Transporter, hosts *Hoster, commands *Commander) *MyCommander {
	g := &MyCommander{
		transports: transports,
		hosts:      hosts,
		commands:   commands,
	}

	return g
}

// activate debug mode
func (my *MyCommander) ActivateDebugMode() {
	my.commands.debugMode = true
}

// add command/s
func (my *MyCommander) AddCommands(commands ...*Command) {
	my.commands.add(commands...)
}

// remove command/s
func (my *MyCommander) RemoveCommands(names ...string) {
	my.commands.remove(names...)
}

// validate transport
func (my MyCommander) ValidTransport(transport string) bool {
	return my.transports.valid(transport)
}

// validate host
func (my MyCommander) ValidHost(host string) bool {
	return my.hosts.valid(host)
}

// validate command
func (my MyCommander) ValidCommand(cmd string) bool {
	return my.commands.valid(cmd)
}

// run cmd
func (my MyCommander) RunCmd(name string, args ...string) (string, error) {
	out, err := my.commands.run(name, args...)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

// TODO
// run cmd
/*func (my MyCommander) RunMultiCmd() string {
	return ""
}*/
