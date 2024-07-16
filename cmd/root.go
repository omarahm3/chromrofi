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
	orderBy string
	limit   int
	rootCmd = &cobra.Command{
		Use:   "chromrofi",
		Short: "chromrofi is a CLI for rofi to browse chrome history",
		Run:   runCommand,
	}
)

func Init() {
	rootCmd.PersistentFlags().StringVarP(&orderBy, "order-by", "o", "last_visit_time", "Property to order by")
	rootCmd.PersistentFlags().IntVarP(&limit, "limit", "l", 10, "Number of results to return")
	if err := rootCmd.Execute(); err != nil {
		message("failed to run chromrofi", &err)
	}
}

func runCommand(cmd *cobra.Command, args []string) {
	c, err := chrome.GetChrome()
	if err != nil {
		message("failed to obtain chrome history database", &err)
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

	if err := openUrl(url.Url); err != nil {
		message("failed to open url in browser", &err)
	}

	os.Exit(0)
}

func openUrl(url string) error {
	return exec.Command("xdg-open", url).Start()
}

func getSelection(args []string) string {
	return strings.Join(args, " ")
}

func message(message string, err *error) {
	if err != nil {
		fmt.Printf("[ERROR]> %s: %s\n", message, *err)
		os.Exit(1)
	}

	fmt.Printf("[INFO]> %s\n", message)
}
