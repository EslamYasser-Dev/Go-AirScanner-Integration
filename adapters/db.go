package adapters

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GormDB struct {
	DB *gorm.DB
}

func NewDBInstance() (*GormDB, error) {
	db, err := gorm.Open(sqlite.Open("cases.db"), &gorm.Config{})
	db.Exec("PRAGMA journal_mode = WAL;")

	if err != nil {
		return nil, err
	}

	return &GormDB{
		DB: db,
	}, nil

}

func (g *GormDB) Migration(models ...any) error {
	return g.DB.AutoMigrate(models...)
}

func (g *GormDB) Close() error {
	sqlDB, err := g.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
