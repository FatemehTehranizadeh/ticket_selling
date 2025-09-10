package workerpool

import (
	"context"
	"fmt"
	"sync"

	// "ticket_selling/internal/app"
	// "ticket_selling/internal/domain"
	// "ticket_selling/internal/repo/mem"
	"ticket_selling/internal/logging"
	"ticket_selling/internal/service"
)

type Pool struct {
	Jobs       chan Job
	Results    chan Result
	WG         *sync.WaitGroup
	NumWorkers int
}

// spawn N workers; each loops: read job → process → send result.
func (p *Pool) Start(ctx context.Context, numWorkers int, ms *service.MarketService) {
	for i := 0; i <= numWorkers; i++ {
		p.WG.Add(1)
		go func() {
			defer p.WG.Done()
			for job := range p.Jobs {
				res := processor(ctx, job, ms)
				p.Results <- res
			}
		}()
	}

}

// send to jobs (respects backpressure if buffered).
func (p *Pool) Submit(ctx context.Context, job ...Job) {
	for _, j := range job {
		p.Jobs <- j
	}
	// close(p.Jobs)
}

func (p *Pool) ListOfResults() <-chan Result {
	for r := range p.Results {
		fmt.Println(r)
	}
	return nil
}

// close jobs, wait for workers, close results.
func (p *Pool) Stop(ctx context.Context) {
	close(p.Jobs)
	p.WG.Wait()
	close(p.Results)
}

func processor(ctx context.Context, job Job, ms *service.MarketService) Result {
	// var es mem.MemEventStore
	// es = mem.MemEventStore{
	// 	Events: ,
	// }

	if job.Type == "buy" {
		o, err := ms.Buy(ctx, job.Payload.UserID, job.Payload.EventID, job.Payload.Quantity)
		if err != nil {
			logging.Sugar.Logger.Errorln(err)
		}
		res := Result{
			JobID:        job.ID,
			Order:        o,
			IsSuccessful: true,
			Err:          err,
			Latency:      0,
		}
		return res

	} else if job.Type == "sell" {
		o, err := ms.Sell(ctx, job.Payload.UserID, job.Payload.EventID, job.Payload.Quantity)
		if err != nil {
			logging.Sugar.Logger.Errorln(err)
		}
		res := Result{
			JobID:        job.ID,
			Order:        o,
			IsSuccessful: true,
			Err:          err,
			Latency:      0,
		}
		return res
	} else {
		logging.Sugar.Logger.Errorln("invalid request")

	}
	return Result{}
}
