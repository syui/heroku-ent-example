package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Users struct {
	ent.Schema
}

func (Users) Fields() []ent.Field {
	return []ent.Field{

		field.String("user").
		NotEmpty().
		Immutable().
		MaxLen(7).
		Unique(),

		field.Int("first").
		Optional(),

		field.Int("draw").
		Optional(),

		field.Time("created_at").
		Optional().
		Default(func() time.Time {
			return time.Now()
		}),

		field.Time("updated_at").
		Optional().
		Default(func() time.Time {
			return time.Now()
		}),

	}
}

func (Users) Edges() []ent.Edge {
	return []ent.Edge{}
}


