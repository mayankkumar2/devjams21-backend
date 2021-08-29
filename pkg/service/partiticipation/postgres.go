package participation

import (
	"context"
	"github.com/GDGVIT/devjams21-backend/pkg/util"

	"github.com/GDGVIT/devjams21-backend/pkg/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}



func (r *repo) CreateParticipation(ctx context.Context, eventId *uuid.UUID, userID *uuid.UUID, teamName string) (*model.Participation, error) {
	var p *model.Participation
	return p, r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		p = &model.Participation{
			Team: &model.Team{
				TeamName: teamName,
				JoinCode: util.RandStringRunes(32),
			},
			Submission: &model.Submission{
				Meta: model.JSON(map[string]interface{}{}),
			},
			EventID:    eventId,
		}
		if err := tx.Create(p).Error; err != nil {
			return err
		}
		teamHost := &model.TeamXUser{
			UserID:     userID,
			TeamID:     p.Team.ID,
			IsLeader:   true,
			IsAccepted: true,
		}
		if err := tx.Create(teamHost).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *repo) DeleteParticipation(ctx context.Context, p *model.Participation) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", p.ID).Delete(&model.Participation{}).Error; err != nil {
			return err
		}
		if err := tx.Where("team_id = ?", p.TeamID).Delete(&model.TeamXUser{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *repo) FindByID(ctx context.Context, id *uuid.UUID) (*model.Participation, error) {
	p := new(model.Participation)
	return p, r.DB.WithContext(ctx).Find(p, "id = ?", id).Error
}

func (r *repo) GetParticipationTeams(ctx context.Context, eventID *uuid.UUID) ([]model.Team, error) {
	var teams = make([]model.Team, 0,100)

	err := r.DB.WithContext(ctx).
		Preload("TeamXUser.User").
		Find(&teams, "id IN (SELECT team_id FROM participations WHERE event_id = ?)", eventID).Error

	if err != nil {
		return nil, err
	}

	return teams, err
}

func (r *repo) IsUserParticipatingInEvent(ctx context.Context, eventId, userId *uuid.UUID) (*int64, error){
	c := new(int64)
	return c, r.DB.WithContext(ctx).
		Table("participations").
		Joins("JOIN teams t on t.id = participations.team_id").
		Joins("JOIN team_x_users txu on t.id = txu.team_id and t.id = txu.team_id").
		Where("user_id = ?  AND event_id = ?", userId, eventId).Count(c).Error
}


