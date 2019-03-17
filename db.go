package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	Init()
	AddRecord(record ListingRecord)
	RecordExists(url string) bool
}

type ListingRecord struct {
	Title string
	Url   string
}

// region In Memory

type InMemDatabase struct {
	items map[string]ListingRecord
}

func (db *InMemDatabase) Init() {
	db.items = make(map[string]ListingRecord)
}

func (db *InMemDatabase) AddRecord(record ListingRecord) {
	db.items[record.Url] = record
}

func (db *InMemDatabase) RecordExists(url string) bool {
	_, exists := db.items[url]
	return exists
}

// endregion

// region Sqlite

type SqliteDatabase struct {
	connectionString string
}

func NewSqliteDatabase(connectionString string) *SqliteDatabase {
	db := new(SqliteDatabase)
	db.connectionString = connectionString
	return db
}

func (db *SqliteDatabase) Init() {
	database, err := sql.Open("sqlite3", db.connectionString)
	if err != nil {
		log.Fatal("Error connecting to db: ", err)
	}

	database.Exec(`CREATE TABLE IF NOT EXISTS "Listings" (
		"Url"	TEXT NOT NULL UNIQUE,
		"Title"	TEXT,
		PRIMARY KEY("Url")
	);`)
}

func (db *SqliteDatabase) AddRecord(record ListingRecord) {
	database, err := sql.Open("sqlite3", db.connectionString)
	if err != nil {
		log.Fatal("Error connecting to db: ", err)
	}

	_, err = database.Exec(`INSERT INTO "Listings" (Url, Title) VALUES (?, ?)`, record.Url, record.Title)
	if err != nil {
		log.Println("Error adding record to db: ", err)
	}
}

func (db *SqliteDatabase) rowExists(query string, args ...interface{}) bool {
	database, err := sql.Open("sqlite3", db.connectionString)
	if err != nil {
		log.Fatal("Error connecting to db: ", err)
	}
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err = database.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("error checking if row exists '%s' %v", args, err)
	}
	return exists
}

func (db *SqliteDatabase) RecordExists(url string) bool {
	return db.rowExists(`SELECT Url From "Listings" WHERE Url=?`, url)
}

// endregion
