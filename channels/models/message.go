package models

import (
	"database/sql"
	"errors"

	"github.com/jinzhu/gorm"
)

type (
	// Message - contains data about single chat message
	Message struct {
		gorm.Model
		Text    string `json:"text"`
		Channel Channel
	}
)

// getMessagesList - returns a list of messages with a given offset and limit
func getMessagesList(db *sql.DB, start, count int) ([]Message, error) {
	return nil, errors.New("Not implemented")
}

// getMessage - returns single message information
func (m *Message) getMessage(db *sql.DB) error {
	return errors.New("Not implemented")
}

// deleteMessage - removes message from DB
func (m *Message) deleteMessage(db *sql.DB) error {
	return errors.New("Not implemented")
}
