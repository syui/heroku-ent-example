package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/dialect"
	"github.com/masseelch/elk"
	"net/url"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}
//func (Group) Fields() []ent.Field {
//	return []ent.Field{
//		field.Int("id").
//			StructTag(`json:"oid,omitempty"`),
//	}
//}

//// Fields of the Blob.
//func (Blob) Fields() []ent.Field {
//	return []ent.Field{
//		field.UUID("id", uuid.UUID{}).
//			Default(uuid.New).
//			StorageKey("oid"),
//	}
//}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
		MaxLen(10).
		NotEmpty().
		Unique().
		Immutable(),
		field.String("name").Unique().
		MaxLen(8).
		SchemaType(map[string]string{
			dialect.Postgres: "varchar(8)",
		}),
		field.Int("age"),
		field.Time("created_at").
		Immutable().
		Default(time.Now),
		field.Time("updated_at").
		Default(time.Now),
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
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type).
			Annotations(elk.Groups("user")),
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{elk.ReadGroups("user")}
}
