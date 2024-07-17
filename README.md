# Chromrofi

Chromrofi is a CLI tool for `rofi` that allows you to browse and open entries from your Chrome history.

## Features

- List Chrome history entries ordered by a specified property.
- Open selected URL in the default browser (not only chrome).

## Support

This utility supports only Chrome and linux. Tested on arch linux (btw) and wayland (btw).

## Installation

From releases. Or if you already have go installed:

```sh
go install github.com/omarahm3/chromrofi@latest
```

## Usage

Run `chromrofi` to list the 10 (by default) most recent entries in your Chrome history:

```sh
❯ chromrofi -h
chromrofi is a CLI for rofi to browse chrome history

Usage:
  chromrofi [flags]

Flags:
  -h, --help              help for chromrofi
  -l, --limit int         Number of results to return (default 10)
  -o, --order-by string   Property to order by (default "last_visit_time")
  -p, --profile string    Chrome profile to use (default "Default")
  -s, --use-search        Google if no results found
```

```sh
rofi -show cr -modi "cr:chromrofi --order-by visit_count --limit 20" -config ~/.config/rofi/config.rasi
rofi -show cr -modi "cr:chromrofi --profile 'Profile 1' --order-by visit_count --limit 20" -config ~/.config/rofi/config.rasi
```
