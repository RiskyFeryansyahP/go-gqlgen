package postgres

import (
	"context"

	"github.com/confus1on/meetmeup/ent"
)

type MeetupsRepository struct {
	Client *ent.Client
}

func NewMeetupsRepo(c *ent.Client) *MeetupsRepository {
	return &MeetupsRepository{Client: c}
}

func (m *MeetupsRepository) GetMeetups(ctx context.Context) ([]*ent.Meetups, error) {
	meetups, err := m.Client.Meetups.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return meetups, nil
}
