package mysql

import (
	"Game/tools"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	DatabaseUrl = tools.Conf.MySQLServer.User + ":" + tools.Conf.MySQLServer.Password +
		"@" + "tcp" +
		"(" + tools.Conf.MySQLServer.Host + ":" + tools.Conf.MySQLServer.Port + ")" +
		"/" + tools.Conf.MySQLServer.Database +
		"?" + "charset=" + tools.Conf.MySQLServer.Charset +
		"&" + "parseTime=" + tools.Conf.MySQLServer.ParseTime +
		"&" + "loc=" + tools.Conf.MySQLServer.Loc
)

func Newdb() (*sql.DB, error) {
	db, err := sql.Open("mysql", DatabaseUrl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	return db, err
}
