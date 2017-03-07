// storage.go
// Copyright (C) 2017 hackliff <xavier.bruhiere@gmail.com>
// Distributed under terms of the MIT license.

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olekukonko/tablewriter"
)

const DB_DRIVER string = "sqlite3"
const CLIPS_TABLE string = "clips"

type Storage struct {
	db    *sql.DB
	table *tablewriter.Table
}

func NewStorage(dbPath string, reset bool) (*Storage, error) {
	if reset {
		log.Println("reseting database")
		os.Remove(dbPath)
	}

	db, err := sql.Open(DB_DRIVER, dbPath)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "CONTENT"})
	table.SetBorder(true)
	table.SetRowLine(true)

	return &Storage{db, table}, err
}

func (s *Storage) Init() error {
	sql_table := `
	CREATE TABLE IF NOT EXISTS clips(
		id TEXT NOT NULL PRIMARY KEY,
		content TEXT NOT NULL,
		created_at DATETIME
	);
	`
	_, err := s.db.Exec(sql_table)
	return err
}

func (s *Storage) Get(rowID int) (string, error) {
	sqlGet := `
	SELECT content FROM clips
	WHERE rowid = ?
	`

	stmt, _ := s.db.Prepare(sqlGet)
	defer stmt.Close()
	var content string
	rows, _ := stmt.Query(rowID)
	for rows.Next() {
		_ = rows.Scan(&content)
		log.Printf("got it: %v\n", content)
	}

	return content, nil
}

func (s *Storage) List(limit int) {
	sqlReadall := fmt.Sprintf(`
	SELECT rowid, id, content FROM %s
	ORDER BY datetime(created_at) DESC
	LIMIT %d
	`, CLIPS_TABLE, limit)

	log.Printf("scaning for clips")
	rows, _ := s.db.Query(sqlReadall)
	defer rows.Close()

	for rows.Next() {
		var clipShortcut int
		item := NewClip()
		_ = rows.Scan(&clipShortcut, &item.Hash, &item.Content)
		s.table.Append([]string{strconv.Itoa(clipShortcut), item.Content})
	}

	s.table.Render()
}

func (s *Storage) SaveIfNew(c *Clip) error {
	sql_add := `
	INSERT OR REPLACE INTO clips(
		id,
		content,
		created_at
	) VALUES(?, ?, CURRENT_TIMESTAMP)
	`

	stmt, _ := s.db.Prepare(sql_add)
	defer stmt.Close()
	_, err := stmt.Exec(c.Hash, c.Content)
	return err
}
