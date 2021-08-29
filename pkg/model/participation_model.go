package model

import "github.com/google/uuid"

type Participation struct {
	BaseModel
	TeamID       *uuid.UUID `json:"team_id"`
	Team         *Team      `json:"team,omitempty"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EventID      *uuid.UUID
	Event        *Event      `json:"event,omitempty"`
	SubmissionID *uuid.UUID  `json:"submission_id"`
	Submission   *Submission `json:"submission,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ChallengeID  *uuid.UUID  `json:"challenge_id"`
	Challenge    *Challenge  `json:"challenge,omitempty"`
	Score 		int64 `json:"-"`
}
