package postgres

import (
	"context"

	"github.com/confus1on/meetmeup/ent"
	"github.com/confus1on/meetmeup/ent/users"
)

type UsersRepository struct {
	Client *ent.Client
}

func NewUsersRepo(c *ent.Client) *UsersRepository {
	return &UsersRepository{Client: c}
}

func (u *UsersRepository) GetUserById(ctx context.Context, ID int) (*ent.Users, error) {
	user, err := u.Client.Users.
		Query().
		Where(users.IDEQ(ID)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return user, err
}
