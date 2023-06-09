package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Store struct {
	config            *Config
	logger            *logrus.Logger
	db                *sql.DB
	eventRepository   *EventRepository
	commentRepository *CommentRepository
	ticketRepository  *TicketRepository
	userRepository    *UserRepository
}

func NewStore(config *Config, logger *logrus.Logger) *Store {
	return &Store{
		config: config,
		logger: logger,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DataBaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	s.logger.Info("created db connection")
	return nil
}

func (s *Store) Close() error {
	if err := s.db.Close(); err != nil {
		return err
	}
	s.logger.Println("close db connection")
	return nil
}

func (s *Store) Event() *EventRepository {
	if s.eventRepository != nil {
		return s.eventRepository
	}
	s.eventRepository = &EventRepository{
		store: s,
	}
	return s.eventRepository
}

func (s *Store) Ticket() *TicketRepository {
	if s.ticketRepository != nil {
		return s.ticketRepository
	}
	s.ticketRepository = &TicketRepository{
		store: s,
	}
	return s.ticketRepository
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
