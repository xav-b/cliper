// Copyright (C) 2017 hackliff <xavier.bruhiere@gmail.com>
// Distributed under terms of the MIT license.

package main

import (
	"crypto/md5"
	"io"
	"log"

	"github.com/atotto/clipboard"
)

type Clip struct {
	// md5 computed from `Content`
	Hash []byte

	// convenient access for UI
	Shortcut int

	// data directly read from the clipboard
	Content string
}

func hash(content string) []byte {
	h := md5.New()
	io.WriteString(h, content)

	return h.Sum(nil)
}

func NewClip() *Clip {
	copied, err := clipboard.ReadAll()
	if err != nil {
		log.Printf("failed to read clipboard: %v\n", err)
	}

	return &Clip{
		Hash:    hash(copied),
		Content: copied,
	}
}

// Copy encapsulates clipboard write logic
func (c *Clip) Copy() error {
	return clipboard.WriteAll(c.Content)
}
