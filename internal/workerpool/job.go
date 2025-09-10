package workerpool

import "context"

 type Job struct {
	ID int `json:"id" sql:"id"`
	Type string `json:"type" sql:"type"`
	Payload Payload `json:"payload" sql:"payload"`
	Ctx context.Context
 }

 type Payload struct {
	UserID int `json:"user_id" sql:"user_id"`
	EventID int `json:"event_id" sql:"event_id"`
	Quantity int `json:"quantity" sql:"quantity"`
 }

