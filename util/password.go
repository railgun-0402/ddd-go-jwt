package util

import "golang.org/x/crypto/bcrypt"

// パスワードはHash化してDBに登録する
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// Hash済みパスワードの一致を確認
func CheckPassword(hashedPassword string, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}
