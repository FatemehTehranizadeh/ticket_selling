package repo

import (
	"context"
	"ticket_selling/internal/domain"
)

// EventStore and OrderStore are interfaces that describe how your program will talk to storage.
// A bridge to database


// EventStore defines what the service needs for events.
type EventStore interface {
	Get(ctx context.Context, eventID int) (*domain.Event, error)
	Save(ctx context.Context, e *domain.Event) error
	List(ctx context.Context) ([]domain.Event, error)
	ReserveSeats(ctx context.Context, eventID int, qty int) error
	ReleaseSeats(ctx context.Context, eventID int, qty int) error
}

// OrderStore defines what the service needs for orders.
type OrderStore interface {
	Create(ctx context.Context, o *domain.Order) error
	Get(ctx context.Context, orderID int) (*domain.Order, error)
	Update(ctx context.Context, o *domain.Order) error
	List(ctx context.Context) ([]domain.Order, error)
}