// Copyright (C) 2017 hackliff <xavier.bruhiere@gmail.com>
// Distributed under terms of the MIT license.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
)

type CommandFunc func(*Storage, *CliOptions) error

var commands = map[string]CommandFunc{
	// print out clipboard history
	"ls": func(storage *Storage, cli *CliOptions) error {
		storage.List()
		return nil
	},

	"cp": func(storage *Storage, cli *CliOptions) error {
		clipShortcut, _ := strconv.Atoi(flag.Arg(1))
		res, _ := storage.Get(clipShortcut)
		return clipboard.WriteAll(res)
	},

	"watch": func(storage *Storage, cli *CliOptions) error {
		log.Printf("watching for clipboard [refresh=%v]\n", cli.Refresh)
		for {
			c := NewClip()
			if err := storage.SaveIfNew(c); err != nil {
				return err
			}

			time.Sleep(cli.Refresh)
		}
	},
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
		fmt.Printf("invalid command: %s\n (watch|ls|cp)", cli.Command)
		os.Exit(0)
	}
	if err := runnable(storage, cli); err != nil {
		fmt.Printf("cliper failed: %v\n", err)
		os.Exit(1)
	}
}
