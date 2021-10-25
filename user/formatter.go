package user

type UserFormatter struct {
	Uuid          string `json:"uuid"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	Token         string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	userFormatter := UserFormatter{
		Uuid:          user.Uuid,
		Email:         user.Email,
		VerifiedEmail: user.VerifiedEmail,
		Picture:       user.Picture,
		Token:         token,
	}

	return userFormatter
}
