package user_list_usecase

import (
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/repository"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/usecase"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *usecase.Usecase {
	userRepo := repository.New(db)
	return usecase.New(userRepo)
}
