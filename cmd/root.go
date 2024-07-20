package cmd

import (
	"context"
	"fmt"

	"github.com/omarahm3/chromrofi/platform/browser"
	"github.com/omarahm3/chromrofi/platform/database"
	"github.com/spf13/cobra"
)

var (
	orderBy    string
	browserArg string
	limit      int
	profile    string
	useSearch  bool
	rootCmd    = &cobra.Command{
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
	rootCmd.PersistentFlags().StringVarP(&browserArg, "browser", "b", "chrome", "Browser to use")
	if err := rootCmd.Execute(); err != nil {
		message("failed to run chromrofi", &err)
	}
}

func runCommand(cmd *cobra.Command, args []string) {
	if !browser.HasBrowser(browserArg) {
		err := fmt.Errorf("unsupported browser: %s", browserArg)
		message(fmt.Sprintf("supported browsers: %s", browser.SupportedBrowsers), &err)
	}

	c, err := browser.GetBrowser(browser.GetBrowserType(browserArg), profile)
	if err != nil {
		message("failed to initialize chrome instance", &err)
	}
	defer c.Close()

	db := database.NewDatabase(fmt.Sprintf("file:%s?mode=ro", c.GetHistoryLocation()))
	db.Connect(context.Background())
	defer db.Close()

	if len(args) > 0 {
		handleSelection(args, db)
	}

	printSelections(db)
}
