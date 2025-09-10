package mem

import (
	"context"
	"sync"
	"ticket_selling/internal/domain"
	"ticket_selling/internal/logging"
)

////A concrete implementations of OrderStore interface, but stored in memory instead of a database.

type MemOrderStore struct {
	Orders map[int]*domain.Order
	Mux    sync.RWMutex
}

func (m *MemOrderStore) Create(ctx context.Context, o *domain.Order) error {
	m.Mux.Lock()
	defer m.Mux.Unlock()
	if o.ID == 0 {
		o.ID = len(m.Orders) + 1
	}
	m.Orders[o.ID] = o
	return nil
}

func (m *MemOrderStore) Get(ctx context.Context, orderID int) (*domain.Order, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()
	o, ok := m.Orders[orderID]
	if !ok {
		logging.Sugar.Logger.Errorln(domain.ErrOrderNotFound)
		return nil, domain.ErrOrderNotFound
	}
	return o, nil
}

func (m *MemOrderStore) Update(ctx context.Context, o *domain.Order) error {
	m.Mux.Lock()
	defer m.Mux.Unlock()
	if _, ok := m.Orders[o.ID]; !ok {
		logging.Sugar.Logger.Errorln(domain.ErrOrderNotFound)
		return domain.ErrOrderNotFound
	}
	m.Orders[o.ID] = o
	return nil
}

func (m *MemOrderStore) List(ctx context.Context) ([]domain.Order, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()
	var ordersSlice []domain.Order
	for _, v := range m.Orders {
		ordersSlice = append(ordersSlice, *v)
	}
	return ordersSlice, nil
}
