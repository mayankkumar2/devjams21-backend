package model

import (
	"time"
)

type Event struct {
	BaseModel
	EventName        string      `json:"event_name" gorm:"type:varchar(100)"`
	Start            time.Time   `json:"start"`
	End              time.Time   `json:"end"`
	RSVPStart        time.Time   `json:"rsvp_start"`
	RSVPEnd          time.Time   `json:"rsvp_end"`
	Meta             JSON        `json:"meta" gorm:"type:json"`
	MemberLimit      uint        `json:"member_limit"`
	MemberLowerLimit uint        `json:"member_lower_limit"`
	Challenge        []Challenge `json:"challenge,omitempty"`
}
