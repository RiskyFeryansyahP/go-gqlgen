package graphql

import (
	"context"
	"log"

	"github.com/confus1on/meetmeup/models"
)

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Meetups(ctx context.Context, filter *models.FilterMeetup, limit *int, offset *int) ([]*models.Meetup, error) {
	var m []*models.Meetup

	meetups, err := r.MeetupsRepo.GetMeetups(ctx, filter, limit, offset)
	if err != nil {
		log.Println("Error when get meetups", err.Error())
	}

	for _, meetup := range meetups {
		meet := &models.Meetup{
			ID:          meetup.ID,
			Name:        meetup.Name,
			Description: meetup.Description,
			UserID:      meetup.UsersID,
		}
		m = append(m, meet)
	}

	return m, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*models.User, error) {
	users, err := r.UsersRepo.GetUserById(ctx, id)
	if err != nil {
		log.Println("Error while get user by ID", err.Error())
		return nil, err
	}

	user := &models.User{
		ID:       users.ID,
		Username: users.Username,
		Email:    users.Email,
	}

	return user, nil
}
