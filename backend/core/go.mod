module github.com/PrikolTech/alpha/backend/core

go 1.24.0

require (
	github.com/Masterminds/squirrel v1.5.4
	github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2 v2.0.0
	github.com/avito-tech/go-transaction-manager/trm/v2 v2.0.0
	github.com/brianvoe/gofakeit/v7 v7.2.1
	github.com/go-faster/errors v0.7.1
	github.com/go-faster/jx v1.1.0
	github.com/google/uuid v1.6.0
	github.com/jackc/pgx/v5 v5.7.2
	github.com/jmoiron/sqlx v1.4.0
	github.com/joho/godotenv v1.5.1
	github.com/ogen-go/ogen v1.10.0
	github.com/samber/lo v1.49.1
	github.com/stretchr/testify v1.10.0
	go.uber.org/mock v0.5.0
	go.uber.org/multierr v1.11.0
	golang.org/x/sync v0.11.0
)

require (
	github.com/avito-tech/go-transaction-manager/drivers/sql/v2 v2.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dlclark/regexp2 v1.11.5 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-faster/yaml v0.4.6 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/exp v0.0.0-20250215185904-eff6e970281f // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	golang.org/x/tools v0.30.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

tool (
	github.com/ogen-go/ogen/cmd/ogen
	go.uber.org/mock/mockgen
)
