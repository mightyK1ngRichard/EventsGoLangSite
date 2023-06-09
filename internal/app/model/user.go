package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	Age               string `json:"age"`
	Address           string `json:"address"`
	ViewingHistory    string `json:"viewing_history"`
	EncryptedPassword string `json:"-"`
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}
	return nil
}

func encryptString(str string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
