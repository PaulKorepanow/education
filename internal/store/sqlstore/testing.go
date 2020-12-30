package sqlstore

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestStore(t *testing.T, databaseURL string) (*SqlStore, func(...string)) {
	t.Helper()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      false,
			LogLevel:      logger.Silent,
		},
	)

	newLogger.LogMode(logger.Silent)

	st, err := gorm.Open(
		postgres.Open(databaseURL),
		&gorm.Config{
			Logger: newLogger,
		},
	)
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
