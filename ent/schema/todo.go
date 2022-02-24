package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

type Todo struct {
    ent.Schema
}

func (Todo) Fields() []ent.Field {
    return []ent.Field{
        field.String("title"),
        field.Bool("done").
            Optional(),
    }
}
