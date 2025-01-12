package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"

	"github.com/PrikolTech/alpha/backend/core/internal/pkg/httpserver"
	"github.com/PrikolTech/alpha/backend/core/internal/pkg/psql"
	"github.com/PrikolTech/alpha/backend/core/internal/transport/http"
	user_create_handler "github.com/PrikolTech/alpha/backend/core/internal/transport/http/user_create"
	user_get_by_id_handler "github.com/PrikolTech/alpha/backend/core/internal/transport/http/user_get_by_id"
	user_list_handler "github.com/PrikolTech/alpha/backend/core/internal/transport/http/user_list"
	user_create_usecase "github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create"
	user_get_by_id_usecase "github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_by_id"
	user_list_usecase "github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := godotenv.Overload()
	if err != nil {
		return 1
	}

	db, err := psql.Connect()
	if err != nil {
		return 1
	}
	defer db.Close()

	userCreateUsecase := user_create_usecase.New(db)
	userGetByIdUsecase := user_get_by_id_usecase.New(db)
	userListUsecase := user_list_usecase.New(db)

	mux := http.New(http.Handlers{
		UserCreate:  user_create_handler.New(userCreateUsecase),
		UserGetById: user_get_by_id_handler.New(userGetByIdUsecase),
		UserList:    user_list_handler.New(userListUsecase),
	})

	server := httpserver.New(mux)
	server.Start(context.TODO())
	<-server.Err()

	return 0
}
