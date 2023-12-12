package migrations

import (
	"flare-tlc/database"
	"flare-tlc/logger"
	"fmt"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"gorm.io/gorm"
)

var Container MigrationContainer = NewMigrationContainer()

type MigrationContainer interface {
	Add(version string, description string, code func(*gorm.DB) error)
	ExecuteAll(db *gorm.DB) error
}

type migration struct {
	version     string
	description string
	code        func(*gorm.DB) error
}

func executeMigration(db *gorm.DB, m migration) error {
	dbMigration := database.Migration{
		Version:     m.version,
		Description: m.description,
		Status:      database.MigrationPending,
		ExecutedAt:  time.Now(),
	}
	err := database.CreateMigration(db, &dbMigration)
	if err != nil {
		return err
	}

	// Execute migration and measure its duration
	start := time.Now()
	execErr := m.code(db)
	end := time.Now()

	var status database.MigrationStatus
	if execErr != nil {
		status = database.MigrationFailed
	} else {
		status = database.MigrationCompleted
	}
	dbMigration.Status = status
	dbMigration.Duration = int((end.Sub(start)).Milliseconds())
	err = database.UpdateMigration(db, &dbMigration)
	if err != nil {
		return fmt.Errorf("error updating migration %s with status %s, error is %w", m.version, status, err)
	}
	if execErr != nil {
		return fmt.Errorf("error executing migration %s, error is %w", m.version, execErr)
	}
	return nil
}

type migrationContainer struct {
	migrations []migration
}

func NewMigrationContainer() MigrationContainer {
	return &migrationContainer{
		migrations: make([]migration, 0, 50),
	}
}

// Adds a migration
//
//	version - Migration version. Migrations are sorted by version and executed in that order.
//	description - Description of the migration
//	code - Migration code, function without parameters
func (mc *migrationContainer) Add(version string, description string, code func(*gorm.DB) error) {
	mc.migrations = append(mc.migrations, migration{
		version:     version,
		description: description,
		code:        code,
	})
}

func (mc *migrationContainer) ExecuteAll(db *gorm.DB) error {
	dbMigrations, err := database.FetchMigrations(db)
	if err != nil {
		return err
	}

	executedVersions := mapset.NewSet[string]()
	for _, m := range dbMigrations {
		if m.Status != database.MigrationCompleted {
			return fmt.Errorf("there is a PENDING or FAILED migration with version: '%s'. Aborting execution of migrations. Problem should be resolved manually", m.Version)
		}
		executedVersions.Add(m.Version)
	}
	currentVersion := "/"
	if len(dbMigrations) > 0 {
		currentVersion = dbMigrations[len(dbMigrations)-1].Version
	}

	sort.Slice(mc.migrations, func(i, j int) bool {
		return mc.migrations[i].version < mc.migrations[j].version
	})
	executedCount := 0
	for _, m := range mc.migrations {
		if !executedVersions.Contains(m.version) {

			logger.Info("Executing migration %s (%s)", m.description, m.version)
			err := executeMigration(db, m)
			if err != nil {
				return err
			}
			currentVersion = m.version
			executedCount++
		}
	}
	logger.Info("Executed %d migrations, current version is %s", executedCount, currentVersion)
	return nil
}
