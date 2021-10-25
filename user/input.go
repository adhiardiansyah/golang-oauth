package user

type RegisterUserInput struct {
	Uuid          string `json:"id" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	VerifiedEmail bool   `json:"verified_email" binding:"required"`
	Picture       string `json:"picture" binding:"required"`
}

type SignInInput struct {
	Email string `json:"email" binding:"required,email"`
}
