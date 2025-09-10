package service

import (
	"context"
	"ticket_selling/internal/domain"
	"ticket_selling/internal/logging"
	"ticket_selling/internal/repo"
	"time"
)

type MarketService struct {
	EventsFromDB repo.EventStore
	OrdersFromDB repo.OrderStore
	//clock
}

func (m *MarketService) Buy(ctx context.Context, userID int, eventID int, qty int) (*domain.Order, error) {
	//TODO: Invalid userID and eventID
	if qty > 0 {
		es := m.EventsFromDB
		eventFromEventStore, err := es.Get(ctx, eventID)
		if err != nil {
			logging.Sugar.Logger.Errorln(domain.ErrEventNotFound)
			return nil, domain.ErrEventNotFound
		}

		if !eventFromEventStore.OnSale {
			logging.Sugar.Logger.Errorln(domain.ErrNotOnSale)
			return nil, domain.ErrNotOnSale
		} else {

			err = eventFromEventStore.Reserve(qty)
			if err != nil {
				logging.Sugar.Logger.Errorln(err)
				return nil, err
			}
			order := domain.Order{
				ID:        userID,
				UserID:    userID,
				EventID:   eventID,
				Quantity:  qty,
				Status:    "Created",
				CreatedAt: time.Now(),
			}
			order.MarkReserved()

			// Simulate payment
			payment := true
			if payment {
				os := m.OrdersFromDB
				order.MarkConfirmed()
				err = os.Create(ctx, &order)
				if err != nil {
					logging.Sugar.Logger.Errorln(err)
					return nil, err
				}
				return &order, nil
			} else {
				order.MarkFailed()
				eventFromEventStore.Release(qty)
				return nil, domain.ErrUnsuccessfulPayment
			}

		}
	} else {
		logging.Sugar.Logger.Errorln(domain.ErrInvalidQuantity)
		return nil, domain.ErrInvalidQuantity
	}

}

/*
Buy flow:

Validate qty > 0. If not → return ErrInvalidQuantity.

Get event from EventStore. If not found → error.

Check event.OnSale. If false → return ErrNotOnSale.

Try event.Reserve(qty) (inside store so it’s safe with lock). If not enough seats → return ErrInsufficientSeats.

Create order with Status = "Reserved".

Simulate payment:

If success → order.MarkConfirmed().

If failure → event.Release(qty) + order.MarkFailed().

Save order in OrderStore.

Return order.
*/

func (m *MarketService) Sell(ctx context.Context, userID int, eventID int, qty int) (*domain.Order, error) {
	if qty > 0 {
		es := m.EventsFromDB
		eventFromEventStore, err := es.Get(ctx, eventID)

		if err != nil {
			logging.Sugar.Logger.Errorln(domain.ErrEventNotFound)
			return nil, domain.ErrEventNotFound
		}

		if !eventFromEventStore.OnSale {
			logging.Sugar.Logger.Errorln(domain.ErrNotOnSale)
			return nil, domain.ErrNotOnSale
		} else {

			err = eventFromEventStore.Release(qty)
			if err != nil {
				return nil, err
			}
			order := domain.Order{
				ID:        0,
				UserID:    userID,
				EventID:   eventID,
				Quantity:  qty,
				Status:    "Created",
				CreatedAt: time.Now(),
			}
			order.MarkConfirmed()

			os := m.OrdersFromDB
			err = os.Create(ctx, &order)
			if err != nil {
				logging.Sugar.Logger.Errorln(err)
				return nil, err
			}
			return &order, nil
		}
	} else {
		logging.Sugar.Logger.Errorln(domain.ErrInvalidQuantity)
		return nil, domain.ErrInvalidQuantity
	}
}

/*
Validate qty > 0.

Get event. If not found or not on sale → error.

Increase available seats with Release(qty).

Create order with Status = "Confirmed".

Save order in store.

Return order.
*/
