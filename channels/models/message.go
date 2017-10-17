package models

import (
	"database/sql"
)

type (
	// Message - contains data about single chat message
	Message struct {
		ID   int    `json:"id"`
		Text string `json:"text"`
	}
)

// GetMessagesList - returns a list of messages with a given offset and limit
func GetMessagesList(db *sql.DB, start, count int) ([]Message, error) {
	rows, err := db.Query(
		"SELECT id, text FROM messages LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := []Message{}

	for rows.Next() {
		var m Message
		if err := rows.Scan(&m.ID, &m.Text); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	return messages, nil
}

// GetMessage - returns single message information
func (m *Message) GetMessage(db *sql.DB) error {
	return db.QueryRow("SELECT text FROM messages WHERE id=$1",
		m.ID).Scan(&m.Text)
}

// CreateMessage - creates a new message
func (m *Message) CreateMessage(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO messages(text) VALUES($1) RETURNING id",
		m.Text).Scan(&m.ID)

	return err
}

// DeleteMessage - removes message from DB
func (m *Message) DeleteMessage(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM messages WHERE id=$1", m.ID)

	return err
}
