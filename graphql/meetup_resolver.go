package graphql

import (
	"context"

	"github.com/confus1on/meetmeup/models"
)

func (r *Resolver) Meetup() MeetupResolver {
	return &meetupResolver{Resolver: r}
}

type meetupResolver struct{ *Resolver }

func (m *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	/*
		* without data loader
		*
		u, err := m.UsersRepo.GetUserById(ctx, obj.UserID)
		if err != nil {
			log.Println("Error when called user", err.Error())
			return nil, err
		}

		user := &models.User{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		}

		return user, nil
	*/

	// with data loader
	return getUserLoader(ctx).Load(obj.UserID)

}
