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
	e := new(model.Event)
	err := r.DB.WithContext(ctx).Transaction(
		func(ex *gorm.DB) error {
			e.EventName = payload.EventName
			e.Start = payload.Start
			e.End = payload.End
			e.RSVPStart = payload.RSVPStart
			e.RSVPEnd = payload.RSVPEnd
			e.Meta = payload.Meta
			e.MemberLimit = payload.MemberLimit
			// Challenges

			err := ex.Create(e).Error
			return err

		},
	)
	if err != nil {
		return nil, err
	}
	return e, nil
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
