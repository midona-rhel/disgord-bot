package disgordbot

import (
	"testing"
)

func TestBot_AddCommand(t *testing.T) {
	b := new(Bot)
	if err := b.AddCommand(
		Command{
			Name:  "One",
			Short: "1",
		},
		Command{
			Name:  "Two",
			Short: "2",
		}); err != nil {
		t.Error(err)
	}

	if err := b.AddCommand(
		Command{
			Name: "One",
		}); err == nil {
		t.Error("Expected name error")
	}

	if err := b.AddCommand(
		Command{
			Short: "1",
		}); err == nil {
		t.Error("Expected shortname error")
	}

	if err := b.AddCommand(
		Command{
			Name:  "Three",
			Short: "3",
		},
		Command{
			Name:  "Three",
			Short: "3",
		}); err == nil {
		t.Error("Expected name error")
	}

	if len(b.commands) != 2 {
		t.Errorf("Expected 2 command in the bot, found %d", len(b.commands))
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("The bot did not panic while trying to add a command while the bot is running")
		}
	}()

	b.running = true
	b.AddCommand(Command{})
}
