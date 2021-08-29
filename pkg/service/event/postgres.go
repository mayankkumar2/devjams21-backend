package event

import (
	"context"

	"github.com/GDGVIT/devjams21-backend/api/schema"
	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func (r *repo) CreateEvent(ctx context.Context, payload *schema.CreateEventRequest) (*model.Event, error) {
	db := r.DB.WithContext(ctx)
	event := &model.Event{
		EventName:        payload.EventName,
		Start:            payload.Start,
		End:              payload.End,
		RSVPStart:        payload.RSVPStart,
		RSVPEnd:          payload.RSVPEnd,
		Meta:             model.JSON(payload.Meta),
		MemberLimit:      payload.MemberLimit,
		MemberLowerLimit: payload.MemberLowerLimit,
	}

	return event, db.Create(&event).Error
}

func (r *repo) GetEvent(ctx context.Context, ID *uuid.UUID) (*model.Event, error) {
	e := new(model.Event)
	err := r.DB.WithContext(ctx).First(e, "id = ?", ID).Error
	if err != nil {
		return nil, err
	}
	return e, err
}
func (r *repo) GetAllEvent(ctx context.Context) ([]model.Event, error) {
	events := make([]model.Event, 0, 100)
	return events, r.DB.WithContext(ctx).Find(&events).Error
}

func (r *repo) GetEventByTeamID(ctx context.Context, teamID *uuid.UUID) (*model.Event, error) {
	e := new(model.Event)
	err := r.DB.WithContext(ctx).
		First(
			e, "id IN (SELECT event_id FROM participations p WHERE p.team_id = ? )", teamID).Error
	if err != nil {
		return nil, err
	}
	return e, err
}

func (r *repo) UpdateEvent(ctx context.Context, payload *schema.UpdateEventRequest) error {
	db := r.DB.WithContext(ctx)
	event := new(model.Event)

	err := db.WithContext(ctx).Table("events").First(event, "id = ?", payload.ID).Error

	if err != nil {
		return err
	}

	event.EventName = payload.EventName
	event.Start = payload.Start
	event.End = payload.End
	event.RSVPStart = payload.RSVPStart
	event.RSVPEnd = payload.RSVPEnd
	event.Meta = model.JSON(payload.Meta)
	event.MemberLimit = payload.MemberLimit
	event.MemberLowerLimit = payload.MemberLowerLimit

	err = db.Save(&event).Error

	return err

}

func (r *repo) DeleteEvent(ctx context.Context, ID *uuid.UUID) error {
	return r.DB.WithContext(ctx).Table("events").Where("id = ?", ID).Delete(ID).Error
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}
