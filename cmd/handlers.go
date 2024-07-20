package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/omarahm3/chromrofi/platform/database"
	"github.com/omarahm3/chromrofi/platform/rofi"
)

func printSelections(db *database.Database) {
	urls, err := db.GetOrderedBy(context.Background(), orderBy, limit)
	if err != nil {
		message("failed to get urls from chrome database", &err)
	}

	fmt.Println(rofi.BuildHistory(urls))
}

func handleSelection(args []string, db *database.Database) {
	title := getSelection(args)
	url, err := db.FindSelection(context.Background(), title)
	if err != nil {
		message("failed to find url in chrome database", &err)
	}

	var u string

	if url == nil && useSearch {
		u = searchUrl(title)
	} else if url != nil {
		u = url.Url
	} else {
		os.Exit(0)
	}

	if err := openUrl(u); err != nil {
		message("failed to open url in browser", &err)
	}

	os.Exit(0)
}

func searchUrl(query string) string {
	return fmt.Sprintf("https://www.google.com/search?q=%s", query)
}

func openUrl(url string) error {
	return exec.Command("xdg-open", url).Start()
}

func getSelection(args []string) string {
	return strings.Join(args, " ")
}

func message(message string, err *error) {
	if err != nil {
		fmt.Printf("[ERROR]> %s (%s)\n", message, *err)
		os.Exit(1)
	}

	fmt.Printf("[INFO]> %s\n", message)
}
