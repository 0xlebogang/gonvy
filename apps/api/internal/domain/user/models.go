package user

import (
	ulid "github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PublicID string  `json:"ulid" gorm:"uniqueIndex;not null;size:26"`
	Email    string  `json:"email" gorm:"uniqueIndex;varchar(255);not null"`
	Name     *string `json:"name" gorm:"type:varchar(255)"`
	Password string  `json:"password" gorm:"type:text;not null"`
}

type UserRequest struct {
	Email    string  `json:"email" binding:"required,email,min=3,max=255"`
	Name     *string `json:"name" binding:"omitempty,min=2,max=255"`
	Password string  `json:"password" binding:"required,min=6"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email,min=3,max=255"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserUpdateRequest struct {
	PublicID *string `json:"ulid"`
	Email    *string `json:"email" binding:"omitempty,email,min=3,max=255"`
	Name     *string `json:"name" binding:"omitempty,min=2,max=255"`
	Password *string `json:"password" binding:"omitempty,min=6"`
}

type UserResponse struct {
	BaseModel gorm.Model
	PublicID  string
	Email     string
	Name      *string
}

func (u *User) AsResponse() *UserResponse {
	return &UserResponse{
		BaseModel: gorm.Model{
			ID:        u.Model.ID,
			CreatedAt: u.Model.CreatedAt,
			UpdatedAt: u.Model.UpdatedAt,
			DeletedAt: u.Model.DeletedAt,
		},
		PublicID: u.PublicID,
		Email:    u.Email,
		Name:     u.Name,
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.PublicID == "" {
		u.PublicID = ulid.Make().String()
	}
	return nil
}
