package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Meetups holds the schema definition for the Meetups entity.
type Meetups struct {
	ent.Schema
}

// Fields of the Meetups.
func (Meetups) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.Text("description"),
		field.Int("usersID"),
	}
}

// Edges of the Meetups.
func (Meetups) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", Users.Type),
	}
}
