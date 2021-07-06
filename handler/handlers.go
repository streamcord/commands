package handler

// CommandHandler represents a command or a subcommand.
type CommandHandler struct {
	Aliases  []string
	Callback func(Context)
	Name     string
}

// GroupHandler represents a group of commands under one root command.
//
// Commands in a group are called with "<group name> <subcommand name>".
type GroupHandler struct {
	Aliases  []string
	Commands map[string]CommandHandler
	Name     string
}

// AddCommand is like Parser.AddCommand but it adds a command to the group rather than directly to the parser.
func (g *GroupHandler) AddCommand(name string, callback func(Context), aliases ...string) CommandHandler {
	handler := CommandHandler{
		Aliases:  aliases,
		Callback: callback,
		Name:     name,
	}

	g.Commands[name] = handler
	for _, a := range aliases {
		g.Commands[a] = handler
	}

	return handler
}
