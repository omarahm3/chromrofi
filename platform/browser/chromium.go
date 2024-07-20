package browser

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

var (
	ErrFailedToGetChromiumPath    = fmt.Errorf("failed to get chromium browser path")
	ErrFailedToGetLocalState      = fmt.Errorf("failed to get local state")
	ErrFailedToLoadProfileHistory = fmt.Errorf("failed to load profile history")
	ErrInvalidProfile             = fmt.Errorf("invalid profile")
)

type ChromiumBrowser struct {
	Profile         string
	historyLocation string
	localState      *LocalState
}

func (c *ChromiumBrowser) Close() error {
	return os.Remove(c.historyLocation)
}

func (c *ChromiumBrowser) GetHistoryLocation() string {
	return c.historyLocation
}

func (c *ChromiumBrowser) GetLocalState() (*LocalState, error) {
	return c.localState, nil
}

func GetChromiumBrowser(profile, directory string) (*ChromiumBrowser, error) {
	cpath, err := getChromiumPath(directory)
	if err != nil {
		return nil, ErrFailedToGetChromiumPath
	}

	localState, err := GetLocalState(cpath)
	if err != nil {
		return nil, ErrFailedToGetLocalState
	}

	if !localState.HasProfile(profile) {
		handleInvalidProfileError(profile, localState)
	}

	key := localState.GetProfileKey(profile)
	history := getHistory(cpath, key)

	tmpHistory, err := createTmpHistory(history)
	if err != nil {
		return nil, ErrFailedToLoadProfileHistory
	}

	return &ChromiumBrowser{
		Profile:         key,
		historyLocation: tmpHistory,
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

func getHistory(cpath, profile string) string {
	return fmt.Sprintf("%s/%s/History", cpath, profile)
}

func getChromiumPath(directory string) (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.config/%s", home, directory), nil
}

func handleInvalidProfileError(profile string, localState *ChromiumLocalState) {
	fmt.Printf("invalid profile '%s', available profiles:\n", profile)

	for i, profile := range localState.Profiles {
		fmt.Printf("%d. %s ('%s')\n", i+1, profile.Name, profile.ID)
	}

	os.Exit(1)
}
