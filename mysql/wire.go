//go:build wireinject
// +build wireinject

package mysql

import (
	"database/sql"
	"github.com/google/wire"
)

func InitDb() (*sql.DB, error) {
	wire.Build(Newdb)
	return nil, nil
}
