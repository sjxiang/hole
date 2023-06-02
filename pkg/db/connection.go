package db

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const RETRY_TIMES = 6

type Config struct {
	Addr     string `env:"HOLE_MYSQL_ADDR"     envDefault:"localhost"`
	Port     string `env:"HOLE_MYSQL_PORT"     envDefault:"3600"`
	User     string `env:"HOLE_MYSQL_USER"     envDefault:"hole_builder"`
	Password string `env:"HOLE_MYSQL_PASSWORD" envDefault:"71De5JllWSetLYU"`
	Database string `env:"HOLE_MYSQL_DATABASE" envDefault:"hole_builder"`
}

func GetConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	return cfg, err
}

func NewDbConnection(cfg *Config, logger *zap.SugaredLogger) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	retries := RETRY_TIMES
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&&multiStatements=true&loc=Local`, 
			cfg.User, 
			cfg.Password, 
			cfg.Addr,
			cfg.Port, 
			cfg.Database)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	for err != nil {
		if logger != nil {
			logger.Errorw("Failed to connect to database, %d", retries)
		}
		if retries > 1 {
			retries--
			time.Sleep(10 * time.Second)
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			continue
		}
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorw("error in connecting db ", "db", cfg, "err", err)
		return nil, err
	}

	// check db connection
	err = sqlDB.Ping()
	if err != nil {
		logger.Errorw("error in connecting db ", "db", cfg, "err", err)
		return nil, err
	}

	logger.Infow("connected with db", "db", cfg)

	return db, err
}
