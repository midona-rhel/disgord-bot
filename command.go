package disgordbot

import "github.com/andersfylling/disgord"

type Command struct {
	F          func(disgord.Session, disgord.MessageCreate, CommandArgs) error
	Permission disgord.PermissionBit
	Nsfw       bool
	Name       string
	Short      string
	Help       string
	module     string
}

type CommandArgs struct {
	Prefix string
	Raw    string
	Args   []string
}
