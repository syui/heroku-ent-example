package schema

import (
	"entgo.io/ent"
	//"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	//"github.com/masseelch/elk"
	"entgo.io/ent/dialect"
	"math/rand"
	//"time"
	//"entgo.io/ent/schema/mixin"
)

type Pet struct {
	ent.Schema
}


func (Pet) Mixin() []ent.Mixin {
    return []ent.Mixin{
        TimeStamp{},
    }
}

func (Pet) Fields() []ent.Field {
	r := rand.Intn(20)
	return []ent.Field{
		field.String("name").
		MaxLen(8).
		SchemaType(map[string]string{
			dialect.Postgres: "varchar(8)",
		}),
		field.Int("age").
		Default(r).NonNegative(),
		field.Int("card").
		Default(r).NonNegative(),
		//Positive(),
		//Annotations(elk.Validation("required,gt=0")),
		//field.Time("created_at").
		//Default(time.Now).SchemaType(map[string]string{
		//	dialect.Postgres: "date",
		//}),
		//field.Time("updated_at").
		//Default(time.Now).SchemaType(map[string]string{
		//	dialect.Postgres: "date",
		//}),
	}
}

func (Pet) Edges() []ent.Edge {
	return []ent.Edge{}
}


