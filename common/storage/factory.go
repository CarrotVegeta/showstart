package storage

import (
	"errors"
	"github.com/CarrotVegeta/showstart/common/storage/mysqldb"
	"github.com/CarrotVegeta/showstart/config"
)

type IStorage interface {
	Open(conf *config.Config) error
	CreateDataBase() error
	DB() interface{}
}

func CreateStorage(dbtype string) (IStorage, error) {
	switch dbtype {
	case "mysqldb":
		db := mysqldb.NewMysqlDB()
		if err := db.Open(config.Conf); err != nil {
			return nil, err
		} else {
			return db, nil
		}
	default:
		return nil, errors.New("无效存储类型")
	}
}
