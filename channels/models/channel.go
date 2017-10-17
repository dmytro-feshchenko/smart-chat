package models

import (
	"github.com/jinzhu/gorm"
)

type (
	// Channel - chat channel model
	Channel struct {
		gorm.Model
		Name string `gorm:"size:255"`
	}
)
