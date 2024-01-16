package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a hash from a password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// ValidatePassword compares a hashed password with a password and returns an error if they don't match
func ValidatePassword(hashedPassword, password string) error {
	//compare hashed password with password
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}
