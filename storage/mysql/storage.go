package storage

import (
	"fmt"

	"github.com/ebikode/eLearning-core/config"
	md "github.com/ebikode/eLearning-core/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MDatabase ...
type MDatabase struct {
	db *gorm.DB
}

// Config ...
type Config struct {
	*config.Config
}

// New ...
func New(config *config.Config) *Config {
	return &Config{config}
}

// InitDB ..
func (config *Config) InitDB() (*MDatabase, error) {

	var err error
	customername := config.Database.User
	password := config.Database.Pass
	dbName := config.Database.Name
	dbHost := config.Database.Host
	dbPort := config.Database.Port
	charset := config.Database.Charset
	// charset := "utf8mb4_unicode_ci" // FOR LOCAL DB using mysql 5.0+
	// charset := "utf8mb4_0900_ai_ci" // FOR LIVE SERVER DB using mysql 8.0+

	// First Connect to create a DB if it doesn't exist
	conHostURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True", customername, password, dbHost, dbPort) //Build connection string
	cdb, err := gorm.Open("mysql", conHostURI)
	defer cdb.Close()

	if err != nil {
		return nil, err
	}

	createQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE %s;", dbName, charset)
	cErr := cdb.Exec(createQuery).Error
	if cErr != nil {
		fmt.Println(cErr.Error())
		return nil, cErr
	} else {

		dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", customername, password, dbHost, dbPort, dbName) //Build connection string
		fmt.Println(dbURI)

		mdb := new(MDatabase)

		mdb.db, err = gorm.Open("mysql", dbURI)
		if err != nil {
			fmt.Print(err)
			return mdb, err
		}

		// Migrating tables to database
		mdb.db.Debug().AutoMigrate(
			&md.ActivityLog{},
			&md.Admin{},
			&md.AppSetting{},
			&md.AuthdDevice{},
			&md.Assessment{},
			&md.Application{},
			&md.Grade{},
			&md.Article{},
			&md.Journal{},
			&md.Schedule{},
			&md.Course{},
			&md.Question{},
			&md.User{},
		) //Database migration

		// defer mdb.db.Close()

		return mdb, nil
	}

}
