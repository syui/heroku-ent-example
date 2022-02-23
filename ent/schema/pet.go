package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/masseelch/elk"
	"entgo.io/ent/dialect"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
		MaxLen(8).
		SchemaType(map[string]string{
			dialect.Postgres: "varchar(8)",
		}),
		field.Int("age").
			Positive().
			Annotations(elk.Validation("required,gt=0")),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("pets").
			Unique().
			Required().
			Annotations(elk.Validation("required")),
	}
}
