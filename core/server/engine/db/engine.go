package db

import (
	"fmt"

	"github.com/jaronnie/ragingcd/core/server/domain/po"

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
	Engine, err = xorm.NewEngine(config.Mapping.DB.Type, generateDataSourceName())
	if err != nil {
		return err
	}
	err = Engine.Ping()
	if err != nil {
		return err
	}
	return nil
}

func SyncTable() error {
	return Engine.Sync(po.CodeHosting{})
}

func generateDataSourceName() string {
	var dataSource string
	switch config.Mapping.DB.Type {
	case SQLite:
		{
			dataSource = config.Mapping.DB.Database + "?cache=shared&mode=rwc"
		}
	case MySQL:
		{
			dataSource = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				config.Mapping.DB.Username, config.Mapping.DB.Password, config.Mapping.DB.Address, config.Mapping.DB.Database)
		}

	}
	return dataSource
}
