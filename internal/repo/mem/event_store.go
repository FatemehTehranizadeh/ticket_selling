package mem

import (
	"context"
	"sync"
	"ticket_selling/internal/domain"
	"ticket_selling/internal/logging"
)

//A concrete implementations of EventStore interface, but stored in memory instead of a database.

type MemEventStore struct {
	Events map[int]*domain.Event
	Mux    sync.RWMutex
}

func (m *MemEventStore) Get(ctx context.Context, eventID int) (*domain.Event, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()
	e, ok := m.Events[eventID]
	if !ok {
		logging.Sugar.Logger.Errorln(domain.ErrEventNotFound)
		return nil, domain.ErrEventNotFound
	}
	return e, nil
}

func (m *MemEventStore) Save(ctx context.Context, e *domain.Event) error {
	m.Mux.Lock()
	defer m.Mux.Unlock()
	m.Events[e.ID] = e
	return nil
}

func (m *MemEventStore) List(ctx context.Context) ([]domain.Event, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()
	var eventsSlice []domain.Event
	for _, v := range m.Events {
		eventsSlice = append(eventsSlice, *v)
	}
	return eventsSlice, nil
}

func (m *MemEventStore) ReserveSeats(ctx context.Context, eventID int, qty int) error {
	m.Mux.Lock()
	defer m.Mux.Unlock()
	e, ok := m.Events[eventID]
	if !ok {
		logging.Sugar.Logger.Errorln(domain.ErrEventNotFound)
		return domain.ErrEventNotFound
	}
	err := e.Reserve(qty)
	if err != nil {
		return err
	}
	return nil
}

func (m *MemEventStore) ReleaseSeats(ctx context.Context, eventID int, qty int) error {
	m.Mux.Lock()
	defer m.Mux.Unlock()
	e, ok := m.Events[eventID]
	if !ok {
		logging.Sugar.Logger.Errorln(domain.ErrEventNotFound)
		return domain.ErrEventNotFound
	}
	err := e.Release(qty)
	if err != nil {
		return err
	}
	return nil
}
