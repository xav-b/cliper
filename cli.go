// cli.go
// Copyright (C) 2017 hackliff <xavier.bruhiere@gmail.com>
// Distributed under terms of the MIT license.

package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	DEFAULT_REFRESH = "1s"
	DEFAULT_DB_PATH = "/tmp/clip.db"
	DEFAULT_COMMAND = "ls"
)

type CliOptions struct {
	Command string

	// file path for database initialization
	DBPath string

	// empty database before starting
	Reset bool

	// human readable refresh rate before updating the clipboard history
	Refresh time.Duration
}

func NewCliOptions() (*CliOptions, error) {
	dbPath := flag.String("db", DEFAULT_DB_PATH, "database fs path")
	reset := flag.Bool("reset", false, "empty database before starting")
	rawRefresh := flag.String("refresh", DEFAULT_REFRESH, "clipboard refresh rate")
	flag.Parse()

	refreshRate, err := time.ParseDuration(*rawRefresh)
	if err != nil {
		return nil, err
	}

	var command string
	if flag.NArg() < 1 {
		command = DEFAULT_COMMAND
	} else {
		command = flag.Arg(0)
	}

	return &CliOptions{
		Command: command,
		DBPath:  *dbPath,
		Reset:   *reset,
		Refresh: refreshRate,
	}, nil
}

func fail(err error) {
	fmt.Printf("failed to parse options: %v\n", err)
	flag.Usage()
	os.Exit(1)
}
