package domain

import (
	"time"
)

type Order struct {
	ID        int       `json:"id" sql:"id"`
	UserID    int       `json:"user_id" sql:"user_id"`
	EventID   int       `json:"event_id" sql:"event_id"`
	Quantity  int       `json:"quantity" sql:"quantity"`
	Status    string    `json:"status" sql:"status"` //Created, Reserved, Confirmed, Failed, Cancelled
	CreatedAt time.Time `json:"created_at" sql:"created_at"`
	FailureReason string  `json:"failure_reason" sql:"failure_reason"`
	//UpdatedAt
}

func (o *Order) MarkReserved() {
	o.Status = "Reserved"
}

func (o *Order) MarkConfirmed() {
	o.Status = "Confirmed"
}

func (o *Order) MarkCancelled() {
	o.Status = "Cancelled"
}

func (o *Order) MarkFailed() {
	o.Status = "Failed"
}
