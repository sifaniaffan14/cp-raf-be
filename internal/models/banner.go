package models

import (
	"gorm.io/gorm"
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
	"strings"
)

type Banner struct {
	Id      string         `json:"id" gorm:"primaryKey;type:char(26)"`
	ImageUrl   string         `json:"image_url"`
	Tittle string         `json:"tittle"`
	Description    string         `json:"description"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate generates ULID before inserting to DB
func (b *Banner) BeforeCreate(tx *gorm.DB) (err error) {
	if b.Id == "" {
		t := time.Now().UTC()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		b.Id = strings.ToLower(ulid.MustNew(ulid.Timestamp(t), entropy).String())
	}
	return
}