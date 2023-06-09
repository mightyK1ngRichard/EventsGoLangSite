package store

import (
	"github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	tmpUser := &model.User{}
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := r.store.db.QueryRow(`INSERT INTO users (email, password) VALUES ($1, $2) RETURNING email, password;`,
		u.Email, u.EncryptedPassword).Scan(&tmpUser.Email, &tmpUser.Password); err != nil {
		return nil, err
	}

	return tmpUser, nil
}

func (r *UserRepository) CheckUser(u *model.User) (*model.User, error) {
	tmpUser := &model.User{}
	if err := r.store.db.QueryRow(`SELECT id, password FROM users WHERE email = $1;`,
		u.Email,
	).Scan(&tmpUser.ID, &tmpUser.Password); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(tmpUser.Password), []byte(u.Password)); err != nil {
		return nil, err
	}

	return tmpUser, nil
}
