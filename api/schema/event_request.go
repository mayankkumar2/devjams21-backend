package schema

import (
	"time"

	"github.com/google/uuid"
)

type CreateEventRequest struct {
	EventName        string                 `json:"event_name"`
	Start            time.Time              `json:"start"`
	End              time.Time              `json:"end"`
	RSVPStart        time.Time              `json:"rsvp_start"`
	RSVPEnd          time.Time              `json:"rsvp_end"`
	Meta             map[string]interface{} `json:"meta"`
	MemberLimit      uint                   `json:"member_limit"`
	MemberLowerLimit uint                   `json:"member_lower_limit"`
}

type UpdateEventRequest struct {
	ID               *uuid.UUID             `json:"event_id" binding:"uuid"`
	EventName        string                 `json:"event_name"`
	Start            time.Time              `json:"start"`
	End              time.Time              `json:"end"`
	RSVPStart        time.Time              `json:"rsvp_start"`
	RSVPEnd          time.Time              `json:"rsvp_end"`
	Meta             map[string]interface{} `json:"meta"`
	MemberLimit      uint                   `json:"member_limit"`
	MemberLowerLimit uint                   `json:"member_lower_limit"`
}

type DeleteEventRequest struct {
	ID *uuid.UUID `json:"event_id"`
}

type GetEventRequest struct {
	ID *uuid.UUID `json:"event_id"`
}
