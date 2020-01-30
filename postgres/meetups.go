package postgres

import (
	"context"

	"github.com/confus1on/meetmeup/ent"
	"github.com/confus1on/meetmeup/ent/meetups"
	"github.com/confus1on/meetmeup/models"
)

type MeetupsRepository struct {
	Client *ent.Client
}

func NewMeetupsRepo(c *ent.Client) *MeetupsRepository {
	return &MeetupsRepository{Client: c}
}

func (m *MeetupsRepository) GetMeetups(ctx context.Context, filter *models.FilterMeetup, limit, offset *int) ([]*ent.Meetups, error) {
	meetupsQuery := m.Client.Meetups.
		Query().
		Order(ent.Asc(meetups.FieldID))

	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			meetupsQuery = meetupsQuery.
				Where(meetups.NameContains(*filter.Name))
		}
	}

	if limit != nil {
		meetupsQuery = meetupsQuery.
			Limit(*limit)
	}

	if offset != nil {
		meetupsQuery = meetupsQuery.
			Offset(*offset)
	}

	meetups, err := meetupsQuery.
		All(ctx)

	if err != nil {
		return nil, err
	}
	return meetups, nil
}

func (m *MeetupsRepository) CreateMeetup(ctx context.Context, meetup *models.Meetup) error {
	_, err := m.Client.Meetups.
		Create().
		SetName(meetup.Name).
		SetDescription(meetup.Description).
		SetUsersID(meetup.UserID).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (m *MeetupsRepository) UpdateMeetup(ctx context.Context, id int, meetup *models.Meetup) error {
	_, err := m.Client.Meetups.
		UpdateOneID(id).
		SetName(meetup.Name).
		SetDescription(meetup.Description).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (m *MeetupsRepository) GetMeetupsForUser(ctx context.Context, id int) ([]*ent.Meetups, error) {
	meetups, err := m.Client.Meetups.
		Query().
		Where(meetups.UsersIDEQ(id)).
		Order(ent.Asc(meetups.FieldID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return meetups, nil
}
