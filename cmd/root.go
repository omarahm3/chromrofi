package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/omarahm3/chromrofi/platform/chrome"
	"github.com/omarahm3/chromrofi/platform/database"
	"github.com/omarahm3/chromrofi/platform/rofi"
	"github.com/spf13/cobra"
)

var (
	orderBy   string
	limit     int
	profile   string
	useSearch bool
	rootCmd   = &cobra.Command{
		Use:   "chromrofi",
		Short: "chromrofi is a CLI for rofi to browse chrome history",
		Run:   runCommand,
	}
)

func Init() {
	rootCmd.PersistentFlags().StringVarP(&orderBy, "order-by", "o", "last_visit_time", "Property to order by")
	rootCmd.PersistentFlags().StringVarP(&profile, "profile", "p", "Default", "Chrome profile to use")
	rootCmd.PersistentFlags().IntVarP(&limit, "limit", "l", 10, "Number of results to return")
	rootCmd.PersistentFlags().BoolVarP(&useSearch, "use-search", "s", false, "Google if no results found")
	if err := rootCmd.Execute(); err != nil {
		message("failed to run chromrofi", &err)
	}
}

func runCommand(cmd *cobra.Command, args []string) {
	c, err := chrome.GetChrome(profile)
	if err != nil {
		message("failed to initialize chrome instance", &err)
	}
	defer c.Close()

	db := database.NewDatabase(fmt.Sprintf("file:%s?mode=ro", c.HistoryLocation))
	db.Connect(context.Background())
	defer db.Close()

	if len(args) > 0 {
		handleSelection(args, db)
	}

	printSelections(db)
}

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
