package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"net/url"
	"github.com/google/uuid"
	"regexp"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"unicode/utf8"
	"errors"
	"entgo.io/ent/schema/mixin"
	//"entgo.io/ent/schema/field"
	//"entgo.io/ent/schema"
	//"entgo.io/ent/schema/edge"
	//"github.com/masseelch/elk"
)

type User struct {
	ent.Schema
}

func MaxRuneCount(maxLen int) func(s string) error {
    return func(s string) error {
        if utf8.RuneCountInString(s) > maxLen {
            return errors.New("この文字列は長すぎます")
        }
        return nil
    }
}

type TimeStamp struct {
	mixin.Schema
}

func (TimeStamp) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}
func (User) Mixin() []ent.Mixin {
    return []ent.Mixin{
        TimeStamp{},
    }
}


func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
		MaxLen(1).
		Unique().
		Immutable(),
		field.String("name").
		NotEmpty().
		Unique().
		MaxLen(8).
		SchemaType(map[string]string{
			dialect.Postgres: "varchar(10)",
		}).
		Annotations(entsql.Annotation{
			Size: 10,
		}).
		Validate(MaxRuneCount(7)).
		Match(regexp.MustCompile("[a-z]*")),
		field.JSON("url", &url.URL{}).
		Optional(),
		field.JSON("strings", []string{}).
		Optional(),
		field.Enum("state").
		Values("on", "off").
		Optional(),
		field.Bool("active").
		Default(false),
		field.Float("rank").
		Optional(),
		field.UUID("uuid", uuid.UUID{}).
		Default(uuid.New),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}


//// Edges of the User.
//func (User) Edges() []ent.Edge {
//	return []ent.Edge{
//		edge.To("pets", Pet.Type).
//			Annotations(elk.Groups("user")),
//	}
//}
//
//// Annotations of the User.
//func (User) Annotations() []schema.Annotation {
//	return []schema.Annotation{elk.ReadGroups("user")}
//}
