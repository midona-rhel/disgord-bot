package disgordbot

import (
	"fmt"

	"github.com/andersfylling/disgord"
)

type RuleMap struct {
	disabledRelations *relationMap
}

// A Relation defines the behaviour between a target and a rule.
type Relation uint

// Relations define the behaviour between the target and the rule.
const (
	RelationUndefined = Relation(iota)
	RelationBlocked
	RelationAllowed
)

// A Rule defines relation between a command/module/service (known as an action)
// and a guild/channel/user/role (known as an origin). A Rule is a 1 to 1
// relation and assigning more than one Action or Origin makes the rule invalid.
type Rule struct {
	Action   interface{}
	Origin   interface{}
	Relation Relation
}

func (r *Rule) isValid() bool {
	switch r.Action.(type) {
	case Command, Service, Module:
	default:
		return false
	}
	switch r.Origin.(type) {
	case *disgord.Guild, *disgord.Channel, *disgord.User, *disgord.Role:
	default:
		return false
	}
	return true
}

func (b *RuleMap) CreateRule(r Rule) error {
	if !r.isValid() {
		return fmt.Errorf("The Origin or Action is of an unsupported type")
	}
}
