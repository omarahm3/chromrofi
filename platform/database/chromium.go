package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/omarahm3/chromrofi/platform/browser"
)

func (d *Database) GetOrderedBy(ctx context.Context, property string, limit int) ([]browser.Url, error) {
	var records []browser.Url
	err := d.Dbx.SelectContext(ctx, &records, fmt.Sprintf("SELECT * FROM urls ORDER BY %s DESC LIMIT %d", property, limit))
	return records, err
}

func (d *Database) FindSelection(ctx context.Context, title string) (*browser.Url, error) {
	var record browser.Url
	err := d.Dbx.GetContext(ctx, &record, "SELECT * FROM urls WHERE title = $1", title)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &record, err
}
