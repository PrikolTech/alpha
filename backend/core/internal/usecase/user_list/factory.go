package user_list_usecase

import (
	"github.com/jmoiron/sqlx"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/repository"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/usecase"
)

func New(db *sqlx.DB) *usecase.Usecase {
	userRepo := repository.New(db)
	return usecase.New(userRepo)
}
