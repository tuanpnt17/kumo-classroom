package initialize

import (
	"fmt"
	"github.com/tuanpnt17/kumo-classroom-be/global"
	"github.com/tuanpnt17/kumo-classroom-be/internal/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func InitPostgres() {
	postgresConfig := global.Config.Postgres
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC"
	dsn = fmt.Sprintf(dsn, postgresConfig.Host, postgresConfig.User, postgresConfig.Password, postgresConfig.DbName, postgresConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		global.Logger.Error("Postgres connection failed", zap.String("error", err.Error()))
		panic(err)
	}
	global.Logger.Debug("Postgres connection successful")
	global.AppDB = db
	setPool()
	migrateTables()
}

func setPool() {
	appDb, err := global.AppDB.DB()
	if err != nil {
		global.Logger.Error("Failed to get DB instance", zap.String("error", err.Error()))
		return
	}
	appDb.SetMaxIdleConns(global.Config.Postgres.MaxIdleConns)
	appDb.SetMaxOpenConns(global.Config.Postgres.MaxOpenConns)
	appDb.SetConnMaxLifetime(time.Duration(global.Config.Postgres.ConnMaxLifetimeInSeconds) * time.Second)
}

func migrateTables() {
	err := global.AppDB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		global.Logger.Error("Failed to migrate tables", zap.String("error", err.Error()))
		return
	}
}
