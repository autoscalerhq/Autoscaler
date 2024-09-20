package bootstrap

import (
	"fmt"
	"sync"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var olapDBLock = &sync.Mutex{}
var olapDbInstance *gorm.DB

// GetOLAPDBInstance returns a singleton instance of a GORM database connection for an OLAP database
func GetOLAPDBInstance() (*gorm.DB, error) {
	if dbInstance == nil {
		olapDBLock.Lock()
		defer olapDBLock.Unlock()
		if dbInstance == nil {
			fmt.Println("Creating single GORM database instance now.")
			dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
			db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
			if err != nil {
				return nil, fmt.Errorf("failed to connect to database: %w", err)
			}
			olapDbInstance = db
		} else {
			fmt.Println("GORM database instance already created.")
		}
	} else {
		fmt.Println("GORM database instance already created.")
	}

	return olapDbInstance, nil
}
