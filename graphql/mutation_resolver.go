package graphql

import (
	"context"
	"errors"
	"log"

	"github.com/confus1on/meetmeup/models"
)

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (m *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {

	if len(input.Name) < 3 {
		return nil, errors.New("Name is not long enough!")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("Description is not long enough!")
	}

	meetups := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      1,
	}

	err := m.MeetupsRepo.CreateMeetup(ctx, meetups)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return meetups, nil
}

func (m *mutationResolver) UpdateMeetup(ctx context.Context, id int, input models.UpdateMeetup) (*models.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("Name is not long enough")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("Description is not long enough")
	}

	meetup := &models.Meetup{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
	}

	err := m.MeetupsRepo.UpdateMeetup(ctx, id, meetup)
	if err != nil {
		log.Println("Error while updating meetup", err.Error())
		return nil, err
	}

	return meetup, nil
}
