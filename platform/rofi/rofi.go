package rofi

import (
	"git.sr.ht/~jcmuller/go-rofi/entry"
	"git.sr.ht/~jcmuller/go-rofi/script"
	"github.com/omarahm3/chromrofi/platform/browser"
)

func BuildHistory(urls []browser.Url) string {
	r := script.New()

	for _, e := range urls {
		r.AddEntries(entry.New(e.Title, entry.WithInfo(e.Url)))
	}

	return r.Build()
}
