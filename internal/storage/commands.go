package storage

import (
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

func (s *Storage) SaveSongToDb(groupName, songName, songText, link, releaseDate string) error {

	path := fmt.Sprintf("%s%s", os.Getenv("SQL_PATH"), "save.sql")

	req, err := s.PrepareSQL(path)
	if err != nil {
		return err
	}

	_, err = req.Exec(groupName, songName, songText, link, releaseDate)
	if err != nil {
		return err
	}

	slog.Info("Song save")

	return nil
}

func (s *Storage) DeleteSongToDb(id string) error {

	path := fmt.Sprintf("%s%s", os.Getenv("SQL_PATH"), "delete.sql")

	req, err := s.PrepareSQL(path)
	if err != nil {
		return err
	}

	_, err = req.Exec(id)
	if err != nil {
		return err
	}

	slog.Info("Song Delete")
	return nil
}

func (s *Storage) UpdateSongFromIdToDb(songId, groupName, songName, songText, link, releaseDate string) error {

	path := fmt.Sprintf("%s%s", os.Getenv("SQL_PATH"), "update.sql")

	req, err := s.PrepareSQL(path)
	if err != nil {
		return err
	}
	if _, err = req.Exec(songId, groupName, songName, songText, link, releaseDate); err != nil {
		return err
	}

	slog.Info("Song update")
	return nil
}
