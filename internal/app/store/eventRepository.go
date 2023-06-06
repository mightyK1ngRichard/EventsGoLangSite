package store

import "github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/model"

type EventRepository struct {
	store *Store
}

func (r *EventRepository) List() ([]*model.Event, error) {
	rows, err := r.store.db.Query(
		`SELECT * FROM events;`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []*model.Event{}
	for rows.Next() {
		e := &model.Event{}
		if err := rows.Scan(&e.ID, &e.Title, &e.Category, &e.Description, &e.StartDatetime, &e.EndDatetime, &e.Price,
			&e.Address, &e.Organizer, &e.Contacts); err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
