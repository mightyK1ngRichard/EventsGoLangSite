package store

import (
	"github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/model"
	"time"
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

	var events = make([]*model.Event, 0)
	for rows.Next() {
		e := &model.Event{}
		if err := rows.Scan(&e.ID, &e.Title, &e.Category, &e.Description, &e.StartDatetime, &e.EndDatetime, &e.Price,
			&e.Address, &e.Organizer, &e.Contacts); err != nil {
			return nil, err
		}
		e.StartDatetime = CorrectDate(e.StartDatetime)
		e.EndDatetime = CorrectDate(e.EndDatetime)
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
		e.StartDatetime = CorrectDate(e.StartDatetime)
		e.EndDatetime = CorrectDate(e.EndDatetime)
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

func (r *EventRepository) EventByTitle(title string) ([]*model.Event, error) {
	rows, err := r.store.db.Query(
		`SELECT * FROM events WHERE LOWER(title) LIKE '%' || $1 || '%';`,
		title,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events = make([]*model.Event, 0)
	for rows.Next() {
		e := &model.Event{}
		if err := rows.Scan(&e.ID, &e.Title, &e.Category, &e.Description, &e.StartDatetime, &e.EndDatetime, &e.Price,
			&e.Address, &e.Organizer, &e.Contacts); err != nil {
			return nil, err
		}
		e.StartDatetime = CorrectDate(e.StartDatetime)
		e.EndDatetime = CorrectDate(e.EndDatetime)
		events = append(events, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
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
		e.CommentDate = CorrectDate(e.CommentDate)
		result = append(result, e)
	}

	if err := comments.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func CorrectDate(date string) string {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return ""
	}
	return t.Format("2 January 2006, Monday 15:04:05 MST")
}
