package browser

import (
	"path/filepath"
	"slices"
)

type BrowserType string

const (
	Chrome BrowserType = "chrome"
	Brave  BrowserType = "brave"
)

var (
	chromeDirectory = "google-chrome"
	braveDirectory  = filepath.Join("BraveSoftware", "Brave-Browser")
)

var SupportedBrowsers = []BrowserType{Chrome, Brave}

type Browser interface {
	Close() error
	GetHistoryLocation() string
	GetLocalState() (*LocalState, error)
}

func GetBrowser(browser BrowserType, profile string) (Browser, error) {
	switch browser {
	case Chrome:
		return GetChromiumBrowser(profile, chromeDirectory)
	case Brave:
		return GetChromiumBrowser(profile, braveDirectory)
	default:
		return nil, nil
	}
}

func GetBrowserType(browser string) BrowserType {
	switch browser {
	case "chrome":
		return Chrome
	case "brave":
		return Brave
	default:
		return ""
	}
}

func HasBrowser(browser string) bool {
	return slices.Contains(SupportedBrowsers, GetBrowserType(browser))
}
