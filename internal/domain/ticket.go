package domain

type Ticket struct {
	EventID  int `json:"event_id" sql:"event_id"`
	Quantity int `json:"quantity" sql:"quantity"`
}
