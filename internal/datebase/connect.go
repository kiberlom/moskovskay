package datebase

import (
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewConnect() {

	// github.com/mattn/go-sqlite3
	_, _ = gorm.Open(sqlite.Open(path.Join(".", "data", "sql.db")), &gorm.Config{})

}
