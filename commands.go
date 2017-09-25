// commands.go

package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

func PrintHistory(storage *Storage, cli *CliOptions) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "CONTENT"})
	table.SetBorder(true)
	table.SetRowLine(true)

	for _, clip := range storage.List(cli.Last) {
		table.Append([]string{strconv.Itoa(clip.Shortcut), clip.Content})
	}

	table.Render()
	return nil
}

func CopyClip(storage *Storage, cli *CliOptions) error {
	clipShortcut, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		return err
	}

	clip, err := storage.Get(clipShortcut)
	if err != nil {
		return err
	}

	log.Printf("copying '%s' to clipboard\n", clip.Content)
	return clip.Copy()
}

func WatchClipboard(storage *Storage, cli *CliOptions) error {
	log.Printf("watching for clipboard [refresh=%v]\n", cli.Refresh)
	for {
		c := NewClip()
		if err := storage.Upsert(c); err != nil {
			return err
		}

		time.Sleep(cli.Refresh)
	}
}
