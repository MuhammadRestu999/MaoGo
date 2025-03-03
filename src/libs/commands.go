package libs

import (
	"strings"
)

var lists []ICommand

func NewCommands(cmd *ICommand) {
	lists = append(lists, *cmd)
}

func GetList() []ICommand {
	return lists
}

func Get(c *NewClientImpl, m *IMessage) {
	prefix := "#"
	for _, cmd := range lists {
		if cmd.Name == strings.ReplaceAll(m.Command, prefix, "") {
			var cmdWithPref bool
			var cmdWithoutPref bool

			if cmd.IsPrefix && strings.HasPrefix(m.Command, prefix) {
				cmdWithPref = true
			} else {
				cmdWithPref = false
			}

			if !cmd.IsPrefix {
				cmdWithoutPref = true
			} else {
				cmdWithoutPref = false
			}

			if !cmdWithPref && !cmdWithoutPref {
				continue
			}

			//Checking
			if cmd.IsOwner && !m.IsOwner {
				continue
			}

			if cmd.IsMedia && m.Media == nil {
				m.Reply("Media Di Butuhkan")
				continue
			}

			if cmd.IsQuerry && m.Querry == "" {
				m.Reply("Querry Di Butuhkan")
				continue
			}

			cmd.Exec(c, m)
		}
	}
}
