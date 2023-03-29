package model

import (
	"time"
)

type Card struct {
	ID           int       `xorm:"not null pk autoincr INT(11) id"`
	CreatedAt    time.Time `xorm:"TIMESTAMP created_at"`
	UpdatedAt    time.Time `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt    time.Time `xorm:"deleted TIMESTAMP deleted_at"`
	OriginalText string    `xorm:"not null TEXT original_text"`
	ParsedText   string    `xorm:"not null TEXT parsed_text"`
	PV           int       `xorm:"not null default 0 INT(11) pv"`
}

func (Card) TableName() string {
	return "card"
}
