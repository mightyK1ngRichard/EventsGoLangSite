package store

import (
	"github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/model"
)

type TicketRepository struct {
	store *Store
}

func (r *TicketRepository) Tickets() ([]*model.Ticket, error) {
	rows, err := r.store.db.Query(
		`SELECT t.id,
         t.price,
         t.purchase_date,
         u.name,
         e.title
		 FROM tickets t
         LEFT JOIN users u on t.user_id = u.id
         LEFT JOIN events e on t.event_id = e.id;`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events = make([]*model.Ticket, 0)
	for rows.Next() {
		e := &model.Ticket{}
		if err := rows.Scan(&e.ID, &e.Price, &e.PurchaseDate, &e.User, &e.Event); err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
