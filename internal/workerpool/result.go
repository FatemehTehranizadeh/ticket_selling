package workerpool

import (
	"ticket_selling/internal/domain"
	"time"
)

type Result struct {
	JobID int           `json:"job_id" sql:"job_id"`
	Order *domain.Order `json:"order" sql:"order"`
	// OrderID int `json:"order_id" sql:"order_id"`
	IsSuccessful bool `json:"is_successful" sql:"is_successful"`
	Err          error
	Latency      time.Duration `json:"latency" sql:"latency"`
}
