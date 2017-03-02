// Copyright (C) 2017 hackliff <xavier.bruhiere@gmail.com>
// Distributed under terms of the MIT license.

package main

import (
	"crypto/md5"
	"io"

	"github.com/atotto/clipboard"
)

type Clip struct {
	Hash    []byte
	Content string
}

func hash(content string) []byte {
	h := md5.New()
	io.WriteString(h, content)

	return h.Sum(nil)
}

func NewClip() *Clip {
	copied, _ := clipboard.ReadAll()

	return &Clip{
		Hash:    hash(copied),
		Content: copied,
	}
}
