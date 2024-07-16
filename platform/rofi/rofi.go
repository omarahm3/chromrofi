package rofi

import (
	"git.sr.ht/~jcmuller/go-rofi-script"
	"github.com/omarahm3/chromrofi/platform/chrome"
)

func BuildHistory(urls []chrome.Url) string {
	r := rofi.New()

	for _, e := range urls {
		r.AddEntries(rofi.NewEntry(e.Title, rofi.WithInfo(e.Url)))
	}

	return r.Build()
}
