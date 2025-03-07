package database

import (
	"testing"
)

func TestDBConnection(t *testing.T) {
	testDB := SetupTestDB(t)
	defer testDB.Cleanup(t)
}
