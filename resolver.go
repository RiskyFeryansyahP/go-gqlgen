package go_gqlgen

import (
	"context"
	"log"

	"github.com/confus1on/meetmeup/models"
	"github.com/confus1on/meetmeup/postgres"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

var meetups = []*models.Meetup{
	{
		ID:          1,
		Name:        "A First Meetup",
		Description: "A Description",
		UserID:      1,
	},
	{
		ID:          2,
		Name:        "A Second Meetup",
		Description: "A Description",
		UserID:      2,
	},
}

var users = []*models.User{
	{
		ID:       1,
		Username: "Risky",
		Email:    "risky@gmail.com",
	},
	{
		ID:       2,
		Username: "Shania",
		Email:    "shania@gmail.com",
	},
}

type Resolver struct {
	MeetupsRepo postgres.MeetupsRepository
	UsersRepo   postgres.UsersRepository
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Meetup() MeetupResolver {
	return &meetupResolver{Resolver: r}
}

func (r *Resolver) User() UserResolver {
	return &userResolver{Resolver: r}
}

type meetupResolver struct{ *Resolver }

func (m *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
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

}

type userResolver struct{ *Resolver }

func (u *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var m []*models.Meetup

	for _, meetup := range meetups {
		if meetup.UserID == obj.ID {
			m = append(m, meetup)
		}
	}

	return m, nil
}

type mutationResolver struct{ *Resolver }

func (m *mutationResolver) CreateMeetup(ctx context.Context, input NewMeetup) (*models.Meetup, error) {
	panic("Impelement Me!")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	var m []*models.Meetup

	meetups, err := r.MeetupsRepo.GetMeetups(ctx)
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
