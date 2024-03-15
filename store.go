package main

import (
	"database/sql"
	"fmt"

	"github.com/eventials/go-tus"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteStore struct {
	tusDB *sql.DB
}

func NewSqliteStore() (tus.Store, error) {
	db, err := sql.Open("sqlite3", "./database/tus.db")

	if err != nil {
		return nil, err
	}

	query := `
		CREATE TABLE IF NOT EXISTS location (
			id INTEGER NOT NULL PRIMARY KEY,
			finger_print TEXT,
			url TEXT,
			UNIQUE (url/*  */)
		);
	`
	_, err = db.Exec(query)

	if err != nil {
		return nil, err
	}

	return &SqliteStore{
		tusDB: db,
	}, nil
}

func (s *SqliteStore) Get(fingerprint string) (string, bool) {
	query := `
		SELECT url FROM location WHERE finger_print = ?
	`
	rows := s.tusDB.QueryRow(query, fingerprint)

	var url string

	err := rows.Scan(&url)

	if err != nil {
		sbox.AddLine(fmt.Sprintf("data reading failed: %s", err.Error()))
		return "", false
	}

	return url, true
}

func (s *SqliteStore) Set(fingerprint, url string) {
	query := `
		INSERT INTO location VALUES(NULL,?,?)
	`
	_, err := s.tusDB.Exec(query, fingerprint, url)

	if err != nil {
		sbox.AddLine(fmt.Sprintf("data insertion failed: %s", err.Error()))
	}
}

func (s *SqliteStore) Delete(fingerprint string) {
	query := `
		DELETE FROM location WHERE finger_print = ?
	`

	_, err := s.tusDB.Exec(query, fingerprint)

	if err != nil {
		sbox.AddLine(fmt.Sprintf("data deletion failed: %s", err.Error()))
	}
}

func (s *SqliteStore) Close() {
	s.tusDB.Close()
}
