package sqlstore

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path"
	"strings"
	"testing"
	"time"
)

func TestStore(t *testing.T, databaseURL, logPath string) (*SqlStore, func(...string)) {
	t.Helper()

	logFile, err := os.OpenFile(path.Join(logPath, "db.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		t.Fatal(err)
	}

	newLogger := logger.New(
		log.New(logFile, "\r\n", log.Lshortfile),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      false,
			LogLevel:      logger.Info,
		},
	)

	db, err := gorm.Open(
		postgres.Open(databaseURL),
		&gorm.Config{
			Logger: newLogger,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	s := NewStore(db)
	return s, func(tables ...string) {
		if len(tables) > 0 {
			if err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))).Error; err != nil {
				t.Fatal(err)
			}
		}
		logFile.Close()
	}
}
