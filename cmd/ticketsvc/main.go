package main

import (
	"context"
	"fmt"
	"sync"
	"ticket_selling/internal/app"
	"ticket_selling/internal/domain"
	"ticket_selling/internal/logging"
	"ticket_selling/internal/repo/mem"
	"ticket_selling/internal/service"
	"ticket_selling/internal/workerpool"
)

/*
Initialize configuration settings (e.g., read from a config file).
Set up your application's dependencies (like services, repositories).
Initialize your worker pool (later).
Start the application, usually by calling the Run() method on the app.
*/

// var memoryEventStore mem.MemEventStore
// memoryEventStore = mem.MemEventStore {

// }

func main() {

	logging.InitLogger()
	sugar := logging.Sugar.Logger
	defer sugar.Sync()

	sugar.Info("Application started")

	event1 := domain.Event{
		ID:             1,
		Venue:          "Concert",
		TotalSeats:     10,
		AvailableSeats: 8,
		OnSale:         true,
	}
	event2 := domain.Event{
		ID:             2,
		Venue:          "Footbal",
		TotalSeats:     5,
		AvailableSeats: 2,
		OnSale:         true,
	}

	eventsList := make(map[int]*domain.Event)
	eventsList[1] = &event1
	eventsList[2] = &event2

	memoryEventStore := mem.MemEventStore{
		Events: eventsList,
		Mux:    sync.RWMutex{},
	}

	// order1 := domain.Order{
	// 	ID:            1,
	// 	UserID:        1,
	// 	EventID:       2,
	// 	Quantity:      1,
	// 	Status:        "Reserved",
	// 	CreatedAt:     time.Now(),
	// 	FailureReason: "successful",
	// }

	ordersList := make(map[int]*domain.Order)
	// ordersList[1] = &order1

	memoryOrderStore := mem.MemOrderStore{
		Orders: ordersList,
		Mux:    sync.RWMutex{},
	}

	ms := service.MarketService{
		EventsFromDB: &memoryEventStore,
		OrdersFromDB: &memoryOrderStore,
	}

	// _, err := ms.Buy(context.Background(), 9, 1, 3)
	// if err != nil {
	// 	fmt.Println("There is an error: ", err)
	// }
	// ms.Buy(context.Background(), 10, 1, 1)
	// // fmt.Println("The order has been submited successfully!")
	// // fmt.Println("Order details:", newOrder)
	// o, _ := ms.OrdersFromDB.List(context.Background())
	// // o, _ := ms.OrdersFromDB.Get(context.Background(),1)
	// fmt.Println(o)

	jobsCh := make(chan workerpool.Job, 5)
	resCh := make(chan workerpool.Result, 5)
	wg := sync.WaitGroup{}
	bkgCtx := context.Background()

	p := workerpool.Pool{
		Jobs:       jobsCh,
		Results:    resCh,
		WG:         &wg,
		NumWorkers: app.NumWorkers,
	}

	job1 := workerpool.Job{
		ID:   1,
		Type: "buy",
		Ctx:  bkgCtx,
		Payload: workerpool.Payload{
			UserID:   1,
			EventID:  1,
			Quantity: 1,
		},
	}

	job2 := workerpool.Job{
		ID:   2,
		Type: "sell",
		Ctx:  bkgCtx,
		Payload: workerpool.Payload{
			UserID:   2,
			EventID:  1,
			Quantity: 3,
		},
	}

	job3 := workerpool.Job{
		ID:   3,
		Type: "buy",
		Ctx:  bkgCtx,
		Payload: workerpool.Payload{
			UserID:   3,
			EventID:  1,
			Quantity: 4,
		},
	}

	p.Submit(bkgCtx, job1, job2, job3)
	p.Start(bkgCtx, p.NumWorkers, &ms)
	p.Stop(bkgCtx)

	o, _ := ms.OrdersFromDB.List(bkgCtx)
	fmt.Println(o)

	p.ListOfResults()

}
