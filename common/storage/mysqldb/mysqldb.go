package mysqldb

import (
	"fmt"
	"github.com/CarrotVegeta/showstart/config"
	"github.com/CarrotVegeta/showstart/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

type MysqlDB struct {
	db  *gorm.DB
	cfg *Config
}
type Config struct {
	UserName     string
	DatabaseName string
	Server       string
	Port         int
	Password     string
}

func NewMysqlDB() *MysqlDB {
	return &MysqlDB{}
}
func (mydb *MysqlDB) DB() interface{} {
	return mydb.db
}
func (mydb *MysqlDB) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mydb.cfg.UserName,
		mydb.cfg.Password,
		mydb.cfg.Server,
		mydb.cfg.Port,
		mydb.cfg.DatabaseName)
}
func (mydb *MysqlDB) Open(conf *config.Config) error {
	mysqlConf := conf.Mysql
	mydb.cfg = &Config{
		UserName:     mysqlConf.UserName,
		DatabaseName: mysqlConf.Name,
		Server:       mysqlConf.Server,
		Port:         mysqlConf.Port,
		Password:     mysqlConf.Password,
	}
	return mydb.NewDB()
}
func (mydb *MysqlDB) NewDB() error {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mydb.DSN(),
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}))
	if err != nil {
		if strings.Contains(err.Error(), "Unknown database") {
			err := mydb.CreateDataBase()
			if err != nil {
				return fmt.Errorf("create databases %s failed :%v", mydb.cfg.DatabaseName, err.Error())
			}
			return mydb.NewDB()
		} else {
			return fmt.Errorf("open database %s failed:%v", mydb.cfg.DatabaseName, err.Error())
		}
	}
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get db.DB failed:%v", err)
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Minute)
	mydb.db = db
	return nil
}
func (mydb *MysqlDB) CreateDataBase() error {
	logger.FileLog.Info("创建数据库:", mydb.DSN)
	dbname := mydb.cfg.DatabaseName
	mydb.cfg.DatabaseName = "mysql"
	err := mydb.NewDB()
	err = mydb.db.Exec(fmt.Sprintf("CREATE DATABASE `%s` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;", dbname)).Error
	if err != nil {
		return fmt.Errorf("failed to create databases %s:%v", mydb.cfg.DatabaseName, err.Error())
	}
	mydb.cfg.DatabaseName = dbname
	return nil
}
