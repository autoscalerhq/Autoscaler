package bootstrap

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dmLock = &sync.Mutex{}
var dbInstance *gorm.DB

// GetDBInstance returns a singleton instance of a GORM database connection for an OLTP database
func GetDBInstance() (*gorm.DB, error) {
	if dbInstance == nil {
		dmLock.Lock()
		defer dmLock.Unlock()
		if dbInstance == nil {
			fmt.Println("Creating single GORM database instance now.")
			dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				return nil, fmt.Errorf("failed to connect to database: %w", err)
			}
			dbInstance = db
		} else {
			fmt.Println("GORM database instance already created.")
		}
	} else {
		fmt.Println("GORM database instance already created.")
	}

	return dbInstance, nil
}
