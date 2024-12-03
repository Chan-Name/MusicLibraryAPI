package storage

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(dbConfig string) (*Storage, error) {

	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s%s", os.Getenv("SQL_PATH"), "library.up.sql")
	dbCreator, err := Opener(path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(dbCreator)
	if err != nil {
		return nil, err
	}

	slog.Info("Library db created")

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) Del() error {

	path := fmt.Sprintf("%s%s", os.Getenv("SQL_PATH"), "library.down.sql")

	dbCreator, err := s.PrepareSQL(path)
	if err != nil {
		return err
	}

	_, err = dbCreator.Exec()
	if err != nil {
		return err
	}

	slog.Info("Library db deleted")

	return nil
}
