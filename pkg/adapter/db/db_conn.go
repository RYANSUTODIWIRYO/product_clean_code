package db

import (
	"errors"
	"log"

	dbConf "github.com/product/internal/config/db"
)

var (
	errorInvalidDbInstance = errors.New("Invalid db instance")
	ErrEmptyRequest = errors.New("request is mandatory")
	instanceDb = make(map[string]DbDriver)
)

const (
	MySql string = "mysql"
)

// DbDriver is object DB
type DbDriver interface {
	Db() interface{}
}

type Transactioner interface {
	Transaction(fc func(tx interface{}) error) error
}

// NewInstanceDb is used to create a new instance DB
func NewInstanceDb(config dbConf.Database) (DbDriver, error) {
	var err error
	var dbName = config.Name

	switch config.Adapter {
	case MySql:
		dbConn, sqlErr := NewGormMySQLDriver(config)
		if sqlErr != nil {
			err = sqlErr
			log.Fatal("Database connection failed.")
		}
		instanceDb[dbName] = dbConn
	default:
		err = errorInvalidDbInstance
	}
	return instanceDb[dbName], err
}