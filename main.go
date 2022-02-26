package main

import (
	"time"
	"t/ent"
	"net/http"
	"math/rand"
	"context"
	"log"
	"os"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"t/ent/ogent"
	"entgo.io/ent/dialect/sql/schema"
)

type User struct {
	user string `json:"user"`
	first int `json:"first"`
	created_at time.Time `json:"created_at"`
}

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func Random(i int) (l int){
	rand.Seed(time.Now().UnixNano())
	l = rand.Intn(20)
	for l == 0 {
		l = rand.Intn(20)
	}
	return
}

type handler struct {
    *ogent.OgentHandler
    client *ent.Client
}

func (h handler) MarkDone(ctx context.Context, params ogent.MarkDoneParams) (ogent.MarkDoneNoContent, error) {
    return ogent.MarkDoneNoContent{}, h.client.Todo.UpdateOneID(params.ID).SetDone(true).Exec(ctx)
}

func (h handler) DrawStart(ctx context.Context, params ogent.DrawStartParams) (ogent.DrawStartNoContent, error) {
    return ogent.DrawStartNoContent{}, h.client.Users.UpdateOneID(params.ID).SetStart(true).Exec(ctx)
}

func (h handler) DrawDone(ctx context.Context, params ogent.DrawDoneParams) (ogent.DrawDoneNoContent, error) {
    return ogent.DrawDoneNoContent{}, h.client.Users.UpdateOneID(params.ID).SetDraw(Random(22)).Exec(ctx)
}

func main() {
	url := os.Getenv("DATABASE_URL") + "?sslmode=require"
	client, err := ent.Open("postgres", url)
	//client, err := Open(url)
	if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	h := handler{
		OgentHandler: ogent.NewOgentHandler(client),
		client:       client,
	}
	srv,err := ogent.NewServer(h)
	//srv,err := ogent.NewServer(ogent.NewOgentHandler(client))
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":" + port, srv); err != nil {
		log.Fatal(err)
	}
}
