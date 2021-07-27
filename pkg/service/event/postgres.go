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
		EventName:   payload.EventName,
		Start:       payload.Start,
		End:         payload.End,
		RSVPStart:   payload.RSVPStart,
		RSVPEnd:     payload.RSVPEnd,
		Meta:        payload.Meta,
		MemberLimit: payload.MemberLimit,
		// Challenges
	}

	return event, db.Create(&event).Error
}

func (r *repo) GetEvent(ctx context.Context, ID *uuid.UUID) (*model.Event, error) {
	e := new(model.Event)
	err := r.DB.WithContext(ctx).First(e, "event_id = ?", ID).Error
	if err != nil {
		return nil, err
	}
	return e, err
}

func (r *repo) UpdateEvent(ctx context.Context, event *model.Event, payload *schema.UpdateEventRequest) error {

	return r.DB.WithContext(ctx).Table("events").Where("id = ?", event.ID).Updates(payload).Error

}

func (r *repo) DeleteEvent(ctx context.Context, ID *uuid.UUID) error {
	return r.DB.WithContext(ctx).Table("events").Delete(ID).Error
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}
