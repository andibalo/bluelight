package user

type FormattedUser struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ImageURL   string `json:"image_url`
}

func FormatUser(user User, token string) FormattedUser {
	formattedUser := FormattedUser{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		ImageURL:   user.AvatarFileName,
		Token:      token,
	}

	return formattedUser
}
