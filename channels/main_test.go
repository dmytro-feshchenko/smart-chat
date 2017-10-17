package main_test

import (
	"log"
	"os"
	"testing"

	main "github.com/technoboom/smart-chat/channels"
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS messages
(
id SERIAL,
text TEXT NOT NULL,
CONSTRAINT messages_pkey PRIMARY KEY (id)
)`

var a main.App

func TestMain(m *testing.M) {
	a = main.App{}

	a.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"),
	)

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM messages")
	a.DB.Exec("ALTER SEQUENCE messages_id_seq RESTART WITH 1")
}
