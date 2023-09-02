package Services

import (
	"auth/models"
)

// Login user login function
func Login(email string, password string) (data interface{}, error string) {

	user := models.GetUserByEmail(email)

	if !CheckPasswordHash(password, user.Password) {
		return "", "Invalid email or password"
	}

	return CreateTokenDataByUser(user)
}

func CreateTokenDataByUser(user models.User) (data interface{}, error string) {
	jwt := Jwt{}
	token, err := jwt.CreateToken(user)
	if err != nil {
		return nil, "Failed to create token"
	}

	return token, "Successfully created token"
}
