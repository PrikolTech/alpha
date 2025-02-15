package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	"github.com/PrikolTech/alpha/backend/core/internal/pkg/httpserver"
	"github.com/PrikolTech/alpha/backend/core/internal/pkg/psql"
	"github.com/PrikolTech/alpha/backend/core/internal/transport/http"
	project_create_handler "github.com/PrikolTech/alpha/backend/core/internal/transport/http/project_create"
	user_create_handler "github.com/PrikolTech/alpha/backend/core/internal/transport/http/user_create"
	user_get_by_id_handler "github.com/PrikolTech/alpha/backend/core/internal/transport/http/user_get_by_id"
	user_list_handler "github.com/PrikolTech/alpha/backend/core/internal/transport/http/user_list"
	project_create_usecase "github.com/PrikolTech/alpha/backend/core/internal/usecase/project_create"
	user_create_usecase "github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create"
	user_get_by_id_usecase "github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_by_id"
	user_list_usecase "github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list"
)

func main() {
	os.Exit(run())
}

func run() int {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	err := godotenv.Overload()
	if err != nil {
		logger.Error("overload env", slog.String("err", err.Error()))
		return 1
	}

	db, err := psql.Connect(ctx)
	if err != nil {
		logger.Error("connect database", slog.String("err", err.Error()))
		return 1
	}
	defer db.Close()

	userCreateUsecase := user_create_usecase.New(db)
	userGetByIdUsecase := user_get_by_id_usecase.New(db)
	userListUsecase := user_list_usecase.New(db)

	projectCreateUsecase := project_create_usecase.New(db)

	mux := http.New(logger, http.Handlers{
		UserCreate:    user_create_handler.New(userCreateUsecase),
		UserGetById:   user_get_by_id_handler.New(userGetByIdUsecase),
		UserList:      user_list_handler.New(userListUsecase),
		ProjectCreate: project_create_handler.New(projectCreateUsecase),
	})

	server := httpserver.New(mux, logger)
	if err := server.Run(ctx); err != nil {
		logger.Error("fail server", slog.String("err", err.Error()))
		return 1
	}

	logger.Info("stop server")
	return 0
}
