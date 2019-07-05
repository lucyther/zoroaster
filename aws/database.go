package aws

import (
	"zoroaster/config"
	"zoroaster/trigger"
)

type IDB interface {
	InitDB(c *config.ZConfiguration)

	LoadTriggersFromDB(table string) ([]*trigger.Trigger, error)

	LogOutcome(table string, outcome *trigger.Outcome, matchId int)

	GetActions(table string, tgId int, userId int) ([]string, error)

	ReadLastBlockProcessed(table string, watOrWac string) int

	SetLastBlockProcessed(table string, blockNo int, watOrWac string)

	LogMatch(table string, match trigger.Match) int
}