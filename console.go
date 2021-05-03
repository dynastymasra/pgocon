package pgocon

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// import golang migrate source file
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

const (
	migrationSourcePath = "file://migration"
	migrationFilePath   = "./migration"
)

// CreateFile create a new migration file
func CreateFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

// CreateMigrationFiles for postgres and put in migration folder
func CreateMigrationFiles(filename string) error {
	if len(filename) == 0 {
		return errors.New("migration filename is not provided")
	}

	timestamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", migrationFilePath, timestamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", migrationFilePath, timestamp, filename)

	if err := CreateFile(upMigrationFilePath); err != nil {
		return err
	}

	if err := CreateFile(downMigrationFilePath); err != nil {
		os.Remove(upMigrationFilePath)
		return err
	}

	return nil
}

// Migration preparation for postgres
func Migration(data *gorm.DB) (*migrate.Migrate, error) {
	db, err := data.DB()
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationSourcePath, "postgres", driver)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// RunMigration run the database migration
func RunMigration(migration *migrate.Migrate) error {
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

// RollbackMigration rollback the database migration
func RollbackMigration(migration *migrate.Migrate) error {
	if err := migration.Steps(-1); err != nil {
		return err
	}
	return nil
}
