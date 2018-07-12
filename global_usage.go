package jhanda

import (
	"fmt"
	"sort"
	"strings"
)

var template = `
Usage:
%s

Commands:
%s
`

func PrintGlobalUsage(flags string, commandSet CommandSet) (string, error) {
	var globalFlags []string
	for _, flag := range strings.Split(flags, "\n") {
		globalFlags = append(globalFlags, fmt.Sprintf("  %s", flag))
	}

	flagsUsage := globalFlags[0]
	for i := 1; i < len(globalFlags); i++ {
		flagsUsage = fmt.Sprintf("%s\n%s", flagsUsage, globalFlags[i])
	}

	var (
		length int
		names  []string
	)

	for name, _ := range commandSet {
		names = append(names, name)
		if len(name) > length {
			length = len(name)
		}
	}

	sort.Strings(names)

	var commands []string
	for _, name := range names {
		command := commandSet[name]
		name = pad(name, " ", length)
		commands = append(commands, fmt.Sprintf("  %s  %s", name, command.Usage().ShortDescription))
	}

	commandsUsage := commands[0]
	for i := 1; i < len(commands); i++ {
		commandsUsage = fmt.Sprintf("%s\n%s", commandsUsage, commands[i])
	}

	return fmt.Sprintf(template, flagsUsage, commandsUsage), nil
}
