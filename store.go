package main

import (
	"database/sql"
	"fmt"

	"github.com/eventials/go-tus"
	_ "github.com/mattn/go-sqlite3"
)

type MemoryStore struct {
	tusDB *sql.DB
	m     map[string]string
}

func NewMemoryStore() (tus.Store, error) {
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

	return &MemoryStore{
		tusDB: db,
		m:     make(map[string]string),
	}, nil
}

func (s *MemoryStore) Get(fingerprint string) (string, bool) {
	query := `
		SELECT url FROM location WHERE fingerprint = ?
	`
	rows, err := s.tusDB.Query(query, fingerprint)

	defer rows.Close()

	var url string

	for rows.Next() {

		err = rows.Scan(&url)

		if err != nil {
			sbox.AddLine(fmt.Sprintf("data reading failed: %s", err.Error()))
			return "", false
		}
	}

	if err != nil {
		sbox.AddLine(fmt.Sprintf("data reading failed: %s", err.Error()))
		return "", false
	}

	return url, true
}

func (s *MemoryStore) Set(fingerprint, url string) {
	query := `
		INSERT INTO location VALUES(NULL,?,?)
	`
	_, err := s.tusDB.Exec(query, fingerprint, url)

	if err != nil {
		sbox.AddLine(fmt.Sprintf("data insertion failed: %s", err.Error()))
	}
}

func (s *MemoryStore) Delete(fingerprint string) {
	query := `
		DELETE FROM location WHERE finger_print = ?
	`

	_, err := s.tusDB.Exec(query, fingerprint)

	if err != nil {
		sbox.AddLine(fmt.Sprintf("data deletion failed: %s", err.Error()))
	}
}

func (s *MemoryStore) Close() {
	s.tusDB.Close()
}
