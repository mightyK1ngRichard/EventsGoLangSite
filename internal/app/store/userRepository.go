package store

import "github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/model"

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
