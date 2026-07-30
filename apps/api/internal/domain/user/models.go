package user

import (
	"time"

	"github.com/0xlebogang/gonvy/api/internal/common"
	ulid "github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type User struct {
	common.BaseModel
	Email    string  `json:"email" gorm:"uniqueIndex;varchar(255);not null" binding:"required,email,min=3,max=255"`
	Name     *string `json:"name" gorm:"type:varchar(255)" binding:"omitempty,min=2,max=255"`
	Password string  `json:"password" gorm:"type:text;not null" binding:"required,min=6"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email,min=3,max=255"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserUpdateRequest struct {
	Email    *string `json:"email" binding:"omitempty,email,min=3,max=255"`
	Name     *string `json:"name" binding:"omitempty,min=2,max=255"`
	Password *string `json:"password" binding:"omitempty,min=6"`
}

type UserResponse struct {
	PublicID  string    `json:"id"`
	Email     string    `json:"email"`
	Name      *string   `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) AsResponse() *UserResponse {
	return &UserResponse{
		PublicID:  u.PublicID,
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.PublicID == "" {
		u.PublicID = ulid.Make().String()
	}
	return nil
}
