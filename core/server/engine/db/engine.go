package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jaronnie/ragingcd/core/config"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

type DriverType = string

const (
	MySQL  DriverType = "mysql"
	SQLite DriverType = "sqlite3"
)

func InitDB() error {
	var err error
	Engine, err = xorm.NewEngine(config.GlobalConfig.DBConfig.Type, generateDataSourceName())
	if err != nil {
		return err
	}
	err = Engine.Ping()
	if err != nil {
		return err
	}
	return nil
}

func generateDataSourceName() string {
	var dataSource string
	switch config.GlobalConfig.DBConfig.Type {
	case SQLite:
		{
			dataSource = config.GlobalConfig.DBConfig.Database + "?cache=shared&mode=rwc"
		}
	case MySQL:
		{
			dataSource = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				config.GlobalConfig.DBConfig.Username, config.GlobalConfig.DBConfig.Password, config.GlobalConfig.DBConfig.Address, config.GlobalConfig.DBConfig.Database)
		}

	}
	return dataSource
}
