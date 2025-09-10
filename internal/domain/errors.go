package domain

import (
	"errors"
)

var (
	ErrNotOnSale           = errors.New("these tickets can not be saled")
	ErrInsufficientSeats   = errors.New("insufficient seats")
	ErrInvalidQuantity     = errors.New("invalid quantity")
	ErrOrderNotFound       = errors.New("order not found")
	ErrEventNotFound       = errors.New("event not found")
	ErrUnsuccessfulPayment = errors.New("payment was not successful")
	// ErrorInvalidNum      = errors.New("you've entered a negative or zero number for reserving")
)
