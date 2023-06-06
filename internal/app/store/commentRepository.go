package store

import "github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/model"

type CommentRepository struct {
	store *Store
}

func (r *EventRepository) commentsOfEvents(id string) ([]*model.Comment, error) {
	rows, err := r.store.db.Query(
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
	defer rows.Close()

	var comments []*model.Comment
	for rows.Next() {
		c := &model.Comment{}
		if err := rows.Scan(&c.UserId, &c.CommentText, &c.CommentDate); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
