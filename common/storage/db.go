package storage

import "gorm.io/gorm"

type DB interface {
	Open() (*gorm.DB, error)
	CreateDatabase() error
}
