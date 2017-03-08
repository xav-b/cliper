// Copyright (C) 2017 hackliff <xavier.bruhiere@gmail.com>
// Distributed under terms of the MIT license.

package main

import (
	"fmt"
	"log"
	"os"
)

type CommandFunc func(*Storage, *CliOptions) error

var commands = map[string]CommandFunc{
	"ls":    PrintHistory,
	"cp":    CopyClip,
	"watch": WatchClipboard,
}

func main() {
	cli, err := NewCliOptions()
	if err != nil {
		fail(err)
	}

	log.Printf("initializing data backend [driver=%s path=%v]\n", DB_DRIVER, cli.DBPath)
	storage, err := NewStorage(cli.DBPath, cli.Reset)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	// create tables
	storage.Init()

	runnable, ok := commands[cli.Command]
	if !ok {
		fmt.Printf("invalid command: %s (pick watch|ls|cp)\n", cli.Command)
		os.Exit(0)
	}
	if err := runnable(storage, cli); err != nil {
		fmt.Printf("cliper failed: %v\n", err)
		os.Exit(1)
	}
}
