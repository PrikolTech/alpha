package user_get_by_id_usecase

import (
	"github.com/jmoiron/sqlx"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_by_id/repository"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_by_id/usecase"
)

func New(db *sqlx.DB) *usecase.Usecase {
	userRepo := repository.New(db)
	return usecase.New(userRepo)
}
