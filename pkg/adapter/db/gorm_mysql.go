package db

import (
	dbConf "github.com/product/internal/config/db"

	_ "github.com/go-sql-driver/mysql" // Initiate mysql driver
	"github.com/jinzhu/gorm"
)

type GormMySQLDriver struct {
	config dbConf.Database
	db *gorm.DB
}

// NewMySQLDriver new object SQL Driver
func NewGormMySQLDriver(config dbConf.Database) (DbDriver, error) {
	dbConn, err := connect(config)
	if err != nil {
		panic("failed to connect database")
	}
	//defer dbConn.Close()

	// Disable table name's pluralization, if set to true, `Product`'s table name will be `product`
	dbConn.SingularTable(true)

	// Enable Logger, show detailed log
	dbConn.LogMode(true)

	return &GormMySQLDriver{
		config: config,
		db:     dbConn,
	}, nil
}

func connect(config dbConf.Database) (*gorm.DB, error) {
	user := config.Username
	password := config.Password
	host := config.Host
	port := config.Port
	dbname := config.Name
	dbConn, err := gorm.Open("mysql", user+":"+password+"@("+host+":"+port+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	return dbConn, err
}

// Db get db instance of gorm
func (m *GormMySQLDriver) Db() interface{} {
	return m.db
}

func (m *GormMySQLDriver) Transaction(fc func(tx interface{}) error) error {
	return m.db.Transaction(func(tx *gorm.DB) error {
		return fc(tx)
	})
}