// Copyright (C) 2017 hackliff <xavier.bruhiere@gmail.com>
// Distributed under terms of the MIT license.

package main

import (
	"crypto/md5"
	"io"
	"log"
	"net/url"

	"github.com/atotto/clipboard"
)

const (
	UNKNOWN_TYPE = iota // 0
	URL_TYPE            // 1
)

func detectType(rawClip string) int {
	// is it an url ?
	if _, err := url.ParseRequestURI(rawClip); err == nil {
		return URL_TYPE
	}

	return UNKNOWN_TYPE
}

type Clip struct {
	// md5 computed from `Content`
	Hash []byte

	// convenient access for UI
	Shortcut int

	// data directly read from the clipboard
	Content string

	// standout useful clips
	Pinned bool

	// try to identify different type of clips, like code snippets, urls, ...
	Type int
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

	detectedType := detectType(copied)

	return &Clip{
		Hash:    hash(copied),
		Content: copied,
		Type:    detectType(copied),
		Pinned:  false,
	}
}

// Copy encapsulates clipboard write logic
func (c *Clip) Copy() error {
	return clipboard.WriteAll(c.Content)
}
