package database

import (
	"errors"
)

func (repository *CheckDatabaseRepository) Check() (bool, error) {
	var res int
	err := repository.database.DB().QueryRow("SELECT 1").Scan(&res)
	if err != nil {
		return false, errors.New("DB connection error")
	}

	if res != 1 {
		return false, errors.New("Unexpected query result")
	}

	return true, nil
}
