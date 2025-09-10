package domain

import "ticket_selling/internal/logging"

type Event struct {
	ID             int    `json:"id" sql:"id"`
	Venue          string `json:"venue" sql:"venue"`
	TotalSeats     int    `json:"total_seats" sql:"total_seats"`
	AvailableSeats int    `json:"available_seats" sql:"available_seats"`
	OnSale         bool   `json:"on_sale" sql:"on_sale"`
}

// Check availability
func (e *Event) CanReserve(n int) error {
	if n <= 0 {
		logging.Sugar.Logger.Errorln(ErrInvalidQuantity)
		return ErrInvalidQuantity
	} else {
		if n <= e.AvailableSeats {
			logging.Sugar.Logger.Infoln("These seats can be reserved!")
		} else {
			logging.Sugar.Logger.Errorln(ErrInsufficientSeats)
			return ErrInsufficientSeats
		}
	}
	return nil
}

// decrease AvailableSeats; fail if insufficient.
func (e *Event) Reserve(n int) error {
	if n <= 0 {
		logging.Sugar.Logger.Errorln(ErrInvalidQuantity)
		return ErrInvalidQuantity
	} else {
		if e.AvailableSeats > 0 {
			if n <= e.AvailableSeats {
				e.AvailableSeats -= n
				logging.Sugar.Logger.Infoln("Congratulations! Your seats have been reserved!")
			} else {
				logging.Sugar.Logger.Errorln(ErrInsufficientSeats)
				return ErrInsufficientSeats
			}
		} else {
			logging.Sugar.Logger.Errorln(ErrInsufficientSeats)
			return ErrInsufficientSeats
		}
	}
	return nil
}

// increase AvailableSeats on cancellation/failed payment.
func (e *Event) Release(n int) error {
	e.AvailableSeats += n
	if e.AvailableSeats > e.TotalSeats {
		e.AvailableSeats = e.TotalSeats
	}
	return nil
}
