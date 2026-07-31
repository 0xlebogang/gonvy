package auth

type UserLogin struct {
	Email    string `json:"email" binding:"required,email,min=3,max=255"`
	Password string `json:"password" binding:"required"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
