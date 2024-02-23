// settings.go
package database

import "database/sql"

// InitSettingsTable initializes the "settings" table.
func InitSettingsTable(db *sql.DB) error {
	createSettingsTableSQL := `
		CREATE TABLE IF NOT EXISTS settings (
			firstRun BOOLEAN DEFAULT false
		);
	`

	_, err := db.Exec(createSettingsTableSQL)
	return err
}

// GetFirstRunSetting retrieves the value of the "firstRun" setting from the "settings" table.
func GetFirstRunSetting(db *sql.DB) (bool, error) {
	query := `SELECT firstRun FROM settings LIMIT 1;`

	var firstRun bool
	err := db.QueryRow(query).Scan(&firstRun)
	return firstRun, err
}

// UpdateFirstRunSetting updates the "firstRun" setting in the "settings" table.
func UpdateFirstRunSetting(db *sql.DB, value bool) error {
	updateSQL := `UPDATE settings SET firstRun = ?;`

	_, err := db.Exec(updateSQL, value)
	return err
}
