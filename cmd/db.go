package cmd

import (
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

func ProvideMysqlDB(dsn string) *gorm.DB {
	mysqlDriver := mysql.New(mysql.Config{
		DSN: dsn,
	})
	db, err := gorm.Open(mysqlDriver, &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}
	return db
}
