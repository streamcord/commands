package handler

import (
	"fmt"
	"strings"

	"github.com/streamcord/commands/internal"
	"github.com/streamcord/commands/types/objects"
)

var (
	// GlobalParser is the Parser struct that stores information like the command/group list and the bot's prefix.
	GlobalParser Parser
)

// Parser is a struct containing information used when parsing messages.
type Parser struct {
	AllowMentionPrefix bool
	Commands           map[string]CommandHandler
	Groups             map[string]GroupHandler
	Prefix             string
}

// ParseMessage will read and parse messages received from the socket and call the matching command, if it exists.
func (p *Parser) ParseMessage(msg objects.Message) {
	// Ignore all messages from bots
	if msg.Author.Bot {
		return
	}

	// Ignore messages without the appropriate prefix
	// We can't assign variables within a condition then use them outside of it, so we have to make a variable declaration before the conditional statement.
	var prefixUsed string
	mentionPrefix := fmt.Sprintf("<@!%s> ", internal.GlobalConf.Discord.ID)

	if strings.HasPrefix(msg.Content, p.Prefix) {
		prefixUsed = p.Prefix
	} else if p.AllowMentionPrefix && strings.HasPrefix(msg.Content, mentionPrefix) {
		prefixUsed = mentionPrefix
	} else {
		return
	}

	// Parse the arguments in the message
	cmd, args := ParseArguments(msg.Content, prefixUsed)

	// Search through regular commands first
	handler := SearchCommand(p.Commands, cmd)
	if handler != nil {
		handler.Callback(NewContext(prefixUsed, msg, args))
		return
	}

	// Failing that, search through the command groups
	// Note: doesn't support nested groups, yet
	group := SearchGroup(p.Groups, cmd)
	if group != nil {
		handler := SearchCommand(group.Commands, args[0])
		if handler != nil {
			handler.Callback(NewContext(prefixUsed, msg, args[1:]))
			return
		}
	}

	// Command not found, so we don't need to do anything
}

// ParseArguments parses the message content and splits away the prefix and command from the rest of the arguments.
func ParseArguments(content string, prefix string) (string, []string) {
	frags := strings.Split(strings.TrimPrefix(content, prefix), " ")
	return frags[0], frags[1:]
}

// CreateParser sets the current global parser, with the provided prefix.
func CreateParser(prefix string) {
	GlobalParser = Parser{
		AllowMentionPrefix: true,
		Commands:           map[string]CommandHandler{},
		Groups:             map[string]GroupHandler{},
		Prefix:             prefix,
	}
}

// AddCommand adds a new command to the global parser.
func (p *Parser) AddCommand(name string, callback func(Context), aliases ...string) CommandHandler {
	handler := CommandHandler{
		Aliases:  aliases,
		Callback: callback,
		Name:     name,
	}

	p.Commands[name] = handler
	for _, a := range aliases {
		p.Commands[a] = handler
	}

	return handler
}

// AddGroup is like AddCommand except it adds a group. To add subcommands to the group, you should call GroupHandler.AddCommand.
func (p *Parser) AddGroup(name string, aliases ...string) GroupHandler {
	handler := GroupHandler{
		Aliases:  aliases,
		Commands: make(map[string]CommandHandler),
		Name:     name,
	}

	p.Groups[name] = handler
	for _, a := range aliases {
		p.Groups[a] = handler
	}

	return handler
}
