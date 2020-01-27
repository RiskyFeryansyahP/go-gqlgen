package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique(),
		field.String("email").
			Unique(),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}
