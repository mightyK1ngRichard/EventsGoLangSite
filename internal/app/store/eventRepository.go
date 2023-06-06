package store

import (
	"github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/model"
)

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

	var events []*model.Event
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

func (r *EventRepository) EventByID(id string) (*model.Event, []*model.Comment, error) {
	e := &model.Event{}
	var comments []*model.Comment
	doneEventData := make(chan error)
	doneCommentsOfEvent := make(chan error)
	var resultEvent error
	var resultComments error

	// Ожидаем заполнение каналов.
	waitForCompletion := func() {
		resultEvent = <-doneEventData
		resultComments = <-doneCommentsOfEvent
	}

	// Асинхронные запросы к бд.
	go func(out chan<- error) {
		var err error
		err = r.store.db.QueryRow(
			`SELECT * FROM events WHERE id = $1;`,
			id,
		).Scan(&e.ID, &e.Title, &e.Category, &e.Description, &e.StartDatetime, &e.EndDatetime, &e.Price,
			&e.Address, &e.Organizer, &e.Contacts)
		if err != nil {
			doneCommentsOfEvent <- err
			return
		}
		doneEventData <- nil
	}(doneEventData)

	go func(out chan<- error) {
		var err error
		comments, err = r.GetCommentsOfPost(id)
		if err != nil {
			doneCommentsOfEvent <- err
			return
		}
		doneCommentsOfEvent <- nil
	}(doneCommentsOfEvent)

	// Ждём выполнение запросов.
	waitForCompletion()

	// Проверяем на наличие ошибок в каналах.
	if resultEvent != nil {
		if resultComments != nil {
			return nil, nil, resultComments
		}
		return nil, nil, resultEvent
	}

	return e, comments, nil
}

func (r *EventRepository) GetCommentsOfPost(id string) ([]*model.Comment, error) {
	comments, err := r.store.db.Query(
		`SELECT 
				u.name,
				c.comment_text,
				c.comment_date
			FROM comments c
				LEFT JOIN events e ON e.id = c.event_id
				LEFT JOIN users u ON c.user_id = u.id
			WHERE e.id = $1;`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer comments.Close()

	var result []*model.Comment
	for comments.Next() {
		e := &model.Comment{}
		if err := comments.Scan(&e.UserId, &e.CommentText, &e.CommentDate); err != nil {
			return nil, err
		}
		result = append(result, e)
	}

	if err := comments.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
