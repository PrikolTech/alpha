package project_create

import (
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/project_create/repository"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/project_create/usecase"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jmoiron/sqlx"

	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
)

func New(db *sqlx.DB) *usecase.Usecase {
	projectRepo := repository.New(db, trmsqlx.DefaultCtxGetter)
	trm := manager.Must(trmsqlx.NewDefaultFactory(db))
	return usecase.New(projectRepo, trm)
}
