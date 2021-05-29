package main

import "golang.org/x/crypto/bcrypt"

func generatePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func comparePassword(password string, passwordHash string) error {
	err := bcrypt.CompareHasAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
