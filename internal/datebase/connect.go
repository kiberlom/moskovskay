package datebase

import (
	"fmt"
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewConnect() error {

	// github.com/mattn/go-sqlite3
	con, err := gorm.Open(sqlite.Open(path.Join("..", "data", "sql.db")), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("Ошибка соединения с бд: %v", err)
	}

	err = migration(con)
	if err != nil {
		return err
	}

	return nil

}
