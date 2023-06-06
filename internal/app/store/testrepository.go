package store

import "github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/model"

type TestRepository struct {
	store *Store
}

func (r *TestRepository) Create(t *model.Test) (*model.Test, error) {
	if err := r.store.db.QueryRow(
		`INSERT INTO test (name, just_info) VALUES ($1, $2) RETURNING id;`,
		t.Name,
		t.Info,
	).Scan(&t.ID); err != nil {
		return nil, err
	}

	return t, nil
}

func (r *TestRepository) FindById(id int) (*model.Test, error) {
	t := &model.Test{}
	if err := r.store.db.QueryRow(
		`SELECT id, name, just_info FROM test WHERE id = $1;`,
		id,
	).Scan(&t.ID, &t.Name, &t.Info); err != nil {
		return nil, err
	}
	return t, nil
}
