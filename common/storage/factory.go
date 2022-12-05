package storage

import (
	"errors"
	"github.com/CarrotVegeta/showstart/conf"
)

type IStorage interface {
	Open(conf *conf.Config) error
	CreateDataBase() error
	DB() interface{}
}

func CreateStorage(dbtype string) (IStorage, error) {
	switch dbtype {
	case "mysqldb":
		db := NewMysqlDB()
		if err := db.Open(conf.Conf); err != nil {
			return nil, err
		} else {
			return db, nil
		}
	default:
		return nil, errors.New("无效存储类型")
	}
}
