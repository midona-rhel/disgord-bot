package disgordbot

import (
	"fmt"

	"github.com/andersfylling/disgord"
)

// Bot represents the instance of the bot running in the
type Bot struct {
	session  disgord.Session
	commands []Command
	running  bool
}

// AddCommand adds a command to the bot, a command needs a unique name and
// shortname to identify it, this function will panic if it is executed while the
// bot is running.
func (b *Bot) AddCommand(cs ...Command) error {
	if b.running {
		panic("Attempted to add Command to Bot while Bot is running")
	}
	var err error
	commands := b.commands
	for _, c := range cs {
		commands, err = testCommand(commands, c)
		if err != nil {
			return err
		}
	}
	b.commands = commands
	return nil
}

func (b *Bot) handleMessageCreate(s disgord.Session, evt *disgord.MessageCreate) {

}

func (b *Bot) computeArgs(s disgord.Session, evt *disgord.MessageCreate) CommandArgs {
	return CommandArgs{}
}

// testCommand tests if there is a name collision for the command and the command slice
func testCommand(cs []Command, c Command) ([]Command, error) {
	for _, command := range cs {
		if command.Name == c.Name {
			return nil, fmt.Errorf("Name collision, the name %s is already defined", c.Short)
		} else if command.Short == c.Short {
			return nil, fmt.Errorf("Name collision, the shortname %s is already defined", c.Short)
		}
	}
	return append(cs, c), nil
}
