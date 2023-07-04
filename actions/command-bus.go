package actions

import (
	"strings"
)

type KookCommand struct {
	IsCommand bool
	Args      []string
	SingleArg string
}

func CommandParse(content string) KookCommand {
	ct := strings.Fields(content)
	if len(ct) != 0 && strings.HasPrefix(ct[0], "/") {
		kc := KookCommand{
			IsCommand: true,
			Args:      ct,
		}
		if len(ct) > 1 {
			kc.SingleArg = strings.SplitN(content, " ", 2)[1]
		}
		return kc
	}
	return KookCommand{IsCommand: false, Args: ct}
}
