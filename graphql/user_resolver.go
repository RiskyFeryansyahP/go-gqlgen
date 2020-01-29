package graphql

import (
	"context"
	"log"

	"github.com/confus1on/meetmeup/models"
)

func (r *Resolver) User() UserResolver {
	return &userResolver{Resolver: r}
}

type userResolver struct{ *Resolver }

func (u *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	meets, err := u.MeetupsRepo.GetMeetupsForUser(ctx, obj.ID)
	if err != nil {
		log.Println("Error while get Meetup for user", err.Error())
		return nil, err
	}

	for _, m := range meets {
		meetup := &models.Meetup{
			ID:          m.ID,
			Name:        m.Name,
			Description: m.Description,
		}

		meetups = append(meetups, meetup)
	}

	return meetups, nil
}
