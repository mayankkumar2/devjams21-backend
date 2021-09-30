package model

import "github.com/google/uuid"

type Participation struct {
	BaseModel
	TeamID       *uuid.UUID  `json:"team_id,omitempty"`
	Team         *Team       `json:"team,omitempty"`
	EventID      *uuid.UUID  `json:"event_id,omitempty"`
	Event        *Event      `json:"event,omitempty"`
	SubmissionID *uuid.UUID  `json:"submission_id,omitempty"`
	Submission   *Submission `json:"submission,omitempty"`
	ChallengeID  *uuid.UUID  `json:"challenge_id,omitempty"`
	Challenge    *Challenge  `json:"challenge,omitempty"`
	Score        int64       `json:"-"`
}
