package schema

import (
	"time"
	"regexp"
	"entgo.io/ent"
	"math/rand"
	"entgo.io/ent/schema/field"
)

type Users struct {
	ent.Schema
}

func Random(i int) (l int){
	rand.Seed(time.Now().UnixNano())
	l = rand.Intn(20)
	for l == 0 {
		l = rand.Intn(20)
	}
	return
}

func (Users) Fields() []ent.Field {
	return []ent.Field{

		field.String("user").
		NotEmpty().
		Immutable().
		MaxLen(7).
		Match(regexp.MustCompile("[a-z]+$")).
		Unique(),

		field.Int("first").
		Immutable().
		Optional().
		Default(Random(20)),

		field.Bool("start").
		Default(false).
		Optional(),

		field.Int("draw").
		Optional().
		Default(Random(20)),

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


