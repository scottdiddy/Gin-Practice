//Package database used gorm to open/read a simple mysql database
//Don't really understand, will learn more in next tutorial and upload in next repo
package database

import (
	"fmt"
	"ginpractice/src/config"
	"time"

	"github.com/allegro/bigcache"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var (
	DB *gorm.DB
	GlobalCache *bigcache.BigCache
)
func Init()  {
	//Connect to DB
	dsn := config.Appconfig.GetString("database.username") + ":" + config.Appconfig.GetString("database.password") +
	"@tcp(" + config.Appconfig.GetString("database.host") + ":" + config.Appconfig.GetString("database.port") +
	")/" + config.Appconfig.GetString("database.name") + "?parseTime=true"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Failed to connect to DB: %w", err))
	}
	GlobalCache, err = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))
	if err != nil {
		panic(fmt.Errorf("Failed to initialize cache: %w", err))

	}
}