package graphql

// go:generate go run github.com/99designs/gqlgen -v

import (
	"github.com/confus1on/meetmeup/postgres"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	MeetupsRepo postgres.MeetupsRepository
	UsersRepo   postgres.UsersRepository
}
