package handler

// SearchCommand searches the provided command mapping for the provided command name.
func SearchCommand(commands map[string]CommandHandler, name string) *CommandHandler {
	// Since this is obviously faster than looping over the map, we're going to try this method first to optimize the function.
	c, ok := commands[name]
	if ok {
		return &c
	}

	// If we got here then we only need to search the aliases.
	for _, c := range commands {
		for _, a := range c.Aliases {
			if a == name {
				return &c
			}
		}
	}

	return nil
}

// SearchGroup is like SearchCommand but instead returns a GroupHandler.
func SearchGroup(groups map[string]GroupHandler, name string) *GroupHandler {
	g, ok := groups[name]
	if ok {
		return &g
	}

	// If we got here then we only need to search the aliases.
	for _, g := range groups {
		for _, a := range g.Aliases {
			if a == name {
				return &g
			}
		}
	}

	return nil
}
