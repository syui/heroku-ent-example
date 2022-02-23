package main

import (
	"context"
	"heroku-ent-example/ent"
	elk "heroku-ent-example/ent/http"
	"fmt"
	"log"
	"os"
	"net/http"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"heroku-ent-example/ent/migrate"
)

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}
func main() {
	url := os.Getenv("DATABASE_URL") + "?sslmode=require"
	c, err := ent.Open("postgres", url)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	defer c.Close()
	ctx := context.Background()
	err = c.Schema.Create(
		ctx, 
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true), 
	)
	users, err := c.User.Query().All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)
	r, l, v := chi.NewRouter(), zap.NewExample(), validator.New()
	//r.Route("/pets", func(r chi.Router) {
	//	elk.NewPetHandler(c, l, v).Mount(r, elk.PetRoutes)
	//})
	r.Route("/users", func(r chi.Router) {
		elk.NewUserHandler(c, l, v).Mount(r, elk.UserRoutes)
	})
	fmt.Println("Server running")
	defer fmt.Println("Server stopped")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":" + port, r); err != nil {
		log.Fatal(err)
	}
}
