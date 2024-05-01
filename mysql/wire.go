//go:build wireinject
// +build wireinject

package mysql

import (
	"database/sql"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitDb() (*sql.DB, error) {
	wire.Build(Newdb)
	return nil, nil
}
func InitOrmDb() (*gorm.DB, error) {
	wire.Build(NewOrmDb)
	return nil, nil
}
