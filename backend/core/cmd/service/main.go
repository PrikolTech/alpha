package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"

	"github.com/PrikolTech/alpha/backend/core/internal/pkg/httpserver"
	"github.com/PrikolTech/alpha/backend/core/internal/pkg/psql"
	"github.com/PrikolTech/alpha/backend/core/internal/transport/http"
	user_create_handler "github.com/PrikolTech/alpha/backend/core/internal/transport/http/user_create"
	user_create_usecase "github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create"
)

func main() {
	os.Exit(run())
}

func run() int {
	godotenv.Overload()

	db, err := psql.Connect()
	if err != nil {
		return 1
	}
	defer db.Close()

	userCreateUsecase := user_create_usecase.New(db)

	mux := http.New(http.Handlers{
		UserCreate: user_create_handler.New(userCreateUsecase),
	})

	server := httpserver.New(mux)
	server.Start(context.TODO())
	<-server.Err()

	return 0
}
