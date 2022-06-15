package mysqldb

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

type MysqlDB struct {
	DB *gorm.DB
}
type Config struct {
	DatabaseName string
	Server       string
	Port         int
	Password     string
}

func NewMysqlDB() *MysqlDB {
	return &MysqlDB{}
}
func (mydb *MysqlDB) DSN(cfg *Config) string {
	return fmt.Sprintf("gorm:gorm@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Server,
		cfg.Port,
		cfg.DatabaseName)
}
func (mydb *MysqlDB) Open(cfg *Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mydb.DSN(cfg),
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}))
	if err != nil {
		if strings.Contains(err.Error(), "Unknown database") {
			err := mydb.CreateDataBase(cfg)
			if err != nil {
				logger.Error("创建数据库失败:%v", err.Error())
				return err
			}
		} else {
			logger.Error("初始化数据库出错，中断启动")
			return err
		}
		this.db, err = gorm.Open("mysql",
			fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
				conf.Key("user").MustString(""),
				conf.Key("pwd").MustString(""),
				conf.Key("server").MustString(""),
				conf.Key("port").MustInt(0),
				conf.Key("name").MustString("")))
		return nil, fmt.Errorf("create gorm db failed:%v", err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get db.DB failed:%v", err)
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Minute)
	return db, nil
}
func (mydb *MysqlDB) CreateDataBase(cfg *Config) error {
	logger.FileLog.Info("创建数据库:", mydb.DSN)
	mydb.cfg.DatabaseName = "mysql"
	cfg := mydb.cfg
	cfg.DatabaseName = "mysql"
	db, err := mydb.Open(cfg)
	err = db.Exec(fmt.Sprintf("CREATE DATABASE `%s` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;", dbname)).Error
	if err != nil {
		return fmt.Errorf("failed to create databases", err.Error())
	}
	return nil
}
