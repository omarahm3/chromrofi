package database

import (
	"context"
	"fmt"

	"github.com/omarahm3/chromrofi/platform/chrome"
)

func (d *Database) GetOrderedBy(ctx context.Context, property string, limit int) ([]chrome.Url, error) {
	var records []chrome.Url
	err := d.Dbx.SelectContext(ctx, &records, fmt.Sprintf("SELECT * FROM urls ORDER BY %s DESC LIMIT %d", property, limit))
	return records, err
}

func (d *Database) FindSelection(ctx context.Context, title string) (chrome.Url, error) {
	var record chrome.Url
	err := d.Dbx.GetContext(ctx, &record, "SELECT * FROM urls WHERE title = $1", title)
	return record, err
}
