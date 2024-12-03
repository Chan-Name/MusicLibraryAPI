package storage

import (
	"database/sql"
	"io"
	"log/slog"
	"os"
)

func Opener(path string) (string, error) {
	sampleInFile, err := os.Open(path)
	if err != nil {
		return "", err
	}

	sample, err := io.ReadAll(sampleInFile)
	if err != nil {
		return "", err
	}

	slog.Info("File with SQL sample open")
	return string(sample), nil
}

func (s *Storage) PrepareSQL(path string) (*sql.Stmt, error) {

	sqlDate, err := Opener(path)
	if err != nil {
		return nil, err
	}

	stmt, err := s.db.Prepare(sqlDate)
	if err != nil {
		return nil, err
	}

	slog.Info("SQL request prepared")
	return stmt, nil
}
