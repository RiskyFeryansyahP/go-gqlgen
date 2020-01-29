// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/confus1on/meetmeup/ent/users"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UsersCreate is the builder for creating a Users entity.
type UsersCreate struct {
	config
	username *string
	email    *string
}

// SetUsername sets the username field.
func (uc *UsersCreate) SetUsername(s string) *UsersCreate {
	uc.username = &s
	return uc
}

// SetEmail sets the email field.
func (uc *UsersCreate) SetEmail(s string) *UsersCreate {
	uc.email = &s
	return uc
}

// Save creates the Users in the database.
func (uc *UsersCreate) Save(ctx context.Context) (*Users, error) {
	if uc.username == nil {
		return nil, errors.New("ent: missing required field \"username\"")
	}
	if uc.email == nil {
		return nil, errors.New("ent: missing required field \"email\"")
	}
	return uc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UsersCreate) SaveX(ctx context.Context) *Users {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UsersCreate) sqlSave(ctx context.Context) (*Users, error) {
	var (
		u     = &Users{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: users.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: users.FieldID,
			},
		}
	)
	if value := uc.username; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: users.FieldUsername,
		})
		u.Username = *value
	}
	if value := uc.email; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: users.FieldEmail,
		})
		u.Email = *value
	}
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}