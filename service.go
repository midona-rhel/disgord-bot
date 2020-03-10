package disgordbot

import "github.com/andersfylling/disgord"

type Service func(disgord.Session, interface{})
