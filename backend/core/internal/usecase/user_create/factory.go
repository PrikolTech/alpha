package user_create_usecase

import (
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create/repository"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create/usecase"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *usecase.Usecase {
	userRepo := repository.New(db)
	trm := manager.Must(trmsqlx.NewDefaultFactory(db))
	return usecase.New(userRepo, trm)
}
