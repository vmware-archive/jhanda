package jhanda

import (
	"fmt"
	"sort"
	"strings"
)

func PrintGlobalUsage(commandSet CommandSet) (string, error) {
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
		commands = append(commands, fmt.Sprintf("%s  %s", name, command.Usage().ShortDescription))
	}

	return strings.Join(commands, "\n"), nil
}
