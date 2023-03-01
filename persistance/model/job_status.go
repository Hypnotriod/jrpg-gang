package model

import (
	"jrpg-gang/engine"
	"time"
)

type JobStatusModel struct {
	Model          `bson:",inline"`
	Email          string                             `bson:"email"`
	IsInProgress   bool                               `bson:"is_in_progress,omitempty"`
	IsComplete     bool                               `bson:"is_complete,omitempty"`
	CompletionTime time.Time                          `bson:"completion_time,omitempty"`
	Code           engine.PlayerJobCode               `bson:"code,omitempty"`
	Countdown      map[engine.PlayerJobCode]time.Time `bson:"countdown"`
}
