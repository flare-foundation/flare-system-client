package database

// Misc other types

type MigrationStatus string

const (
	MigrationPending   MigrationStatus = "PENDING"
	MigrationCompleted MigrationStatus = "COMPLETED"
	MigrationFailed    MigrationStatus = "FAILED"
)
