package sqlstore

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseURL string) (*SqlStore, func(...string)) {
	t.Helper()

	st, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	s := NewStore(st)
	return s, func(tables ...string) {
		if len(tables) > 0 {
			if err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))).Error; err != nil {
				t.Fatal(err)
			}
		}
	}
}
