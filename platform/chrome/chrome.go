package chrome

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

type Url struct {
	ID            int    `db:"id"`
	Title         string `db:"title"`
	Url           string `db:"url"`
	VisitCount    int    `db:"visit_count"`
	TypedCount    int    `db:"typed_count"`
	LastVisitTime int64  `db:"last_visit_time"`
	Hidden        bool   `db:"hidden"`
}

type Chrome struct {
	HistoryLocation string
}

func (c *Chrome) Close() error {
	return os.Remove(c.HistoryLocation)
}

func GetChrome() (*Chrome, error) {
	history, err := getHistory("Default")
	if err != nil {
		return nil, err
	}

	tmpHistory, err := createTmpHistory(history)
	if err != nil {
		return nil, err
	}

	return &Chrome{
		HistoryLocation: tmpHistory,
	}, nil
}

func createTmpHistory(original string) (string, error) {
	tmpPath := filepath.Join(os.TempDir(), "chromrofi-History.db")
	source, err := os.Open(original)
	if err != nil {
		return "", err
	}
	defer source.Close()

	dest, err := os.Create(tmpPath)
	if err != nil {
		return "", err
	}
	defer dest.Close()

	if _, err := io.Copy(dest, source); err != nil {
		return "", err
	}

	return tmpPath, nil
}

func getHistory(profile string) (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.config/google-chrome/%s/History", home, profile), nil
}
