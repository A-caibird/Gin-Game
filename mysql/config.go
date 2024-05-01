package mysql

import (
	"Game/tools"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	OrmUrl = tools.Conf.MySQLServer.User + ":" +
		tools.Conf.MySQLServer.Password + "@" +
		"tcp" + "(" + tools.Conf.MySQLServer.Host + ":" + tools.Conf.MySQLServer.Port + ")" +
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

func NewOrmDb() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       OrmUrl,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	return db, err
}
