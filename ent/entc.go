//go:build ignore

package main

import (
    "log"

    "entgo.io/contrib/entoas"
    "entgo.io/ent/entc"
    "entgo.io/ent/entc/gen"
    "github.com/ariga/ogent"
    "github.com/ogen-go/ogen"
)

func main() {
    spec := new(ogen.Spec)

    //oas, err := entoas.NewExtension(entoas.Spec(spec))
    oas, err := entoas.NewExtension(
        entoas.Spec(spec),
        entoas.Mutations(func(_ *gen.Graph, spec *ogen.Spec) error {
            spec.AddPathItem("/todos/{id}/done", ogen.NewPathItem().
                SetDescription("Mark an item as done").
                SetPatch(ogen.NewOperation().
                    SetOperationID("markDone").
                    SetSummary("Marks a todo item as done.").
                    AddTags("Todo").
                    AddResponse("204", ogen.NewResponse().SetDescription("Item marked as done")),
                ).
                AddParameters(ogen.NewParameter().
                    InPath().
                    SetName("id").
                    SetRequired(true).
                    SetSchema(ogen.Int()),
                ),
            )
            return nil
        }),

								entoas.Mutations(func(_ *gen.Graph, spec *ogen.Spec) error {
									spec.AddPathItem("/users/{id}/start", ogen.NewPathItem().
									SetDescription("Start an draw as done").
									SetPatch(ogen.NewOperation().
									SetOperationID("drawStart").
									SetSummary("Draws a card item as done.").
									AddTags("Users").
                    AddResponse("204", ogen.NewResponse().SetDescription("Item marked as done")),
                ).
                AddParameters(ogen.NewParameter().
                    InPath().
                    SetName("id").
                    SetRequired(true).
                    SetSchema(ogen.Int()),
                ),
            )
            return nil
        }),

								entoas.Mutations(func(_ *gen.Graph, spec *ogen.Spec) error {
									spec.AddPathItem("/users/{id}/d", ogen.NewPathItem().
									SetDescription("Start an draw as done").
									SetPatch(ogen.NewOperation().
									SetOperationID("drawDone").
									SetSummary("Draws a card item as done.").
									AddTags("Users").
                    AddResponse("204", ogen.NewResponse().SetDescription("Item marked as done")),
                ).
                AddParameters(ogen.NewParameter().
                    InPath().
                    SetName("id").
                    SetRequired(true).
                    SetSchema(ogen.Int()),
                ),
            )
            return nil
        }),

    )
    if err != nil {
        log.Fatalf("creating entoas extension: %v", err)
    }
    ogent, err := ogent.NewExtension(spec)
    if err != nil {
        log.Fatalf("creating ogent extension: %v", err)
    }
    err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ogent, oas))
    if err != nil {
        log.Fatalf("running ent codegen: %v", err)
    }
}
