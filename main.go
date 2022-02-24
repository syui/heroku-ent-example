package main

import (
	"time"
	"t/ent"
	"net/http"

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
	srv,err := ogent.NewServer(ogent.NewOgentHandler(client))
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":" + port, srv); err != nil {
		log.Fatal(err)
	}
}
