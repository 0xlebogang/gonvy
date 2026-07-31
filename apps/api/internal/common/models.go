package common

import (
	"time"
)

type BaseModel struct {
	ID        uint      `json:"-" gorm:"primaryKey"`
	PublicID  string    `json:"id" gorm:"uniqueIndex;not null;size:26"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreationTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
