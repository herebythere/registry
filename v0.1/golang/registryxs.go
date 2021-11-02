package registryxs

import (
	"context"
	// "errors"
	"time"
	"fmt"

	qlx "github.com/herebythere/queuelx/v0.1/golang"
)

type SnowprintDetails struct {
	Address	string	`json:"address"`
	Port		string	`json:"port"`
	Kind string `json:"kind"`
	Service string `json:"service"`
	Snowprint	string `json:"snowprint"`
	Timestamp	int64	`json:"timestamp`
}

// service / snowprint / snowprint details
type SnowprintMap = map[string]Snowprints
type SnowprintServiceMap = map[string]SnowprintMap

type Snowprints struct {
	Routers SnowprintMap `json:"routers"`
	Broadcasters SnowprintServiceMap `json:"broadcasters"`
	Receivers SnowprintServiceMap `json:"receivers"`
}

type RegisterCallback func(
	details *SnowprintDetails,
	cancelCallback *context.CancelFunc,
	err error,
) error

type Register struct {
	cleanupInterval int64
	queueInterval int64
	routerService string
	snowprints SnowprintServiceMap
	cancelCallback *context.CancelFunc
	queue *qlx.Queue
}


const (
	broadcaster = "broadcaster"
	reciever = "reciever"
)


func NewRegistry(
	cacheAddress string,
	identifier string,
	routerService string,
	cleanupInterval int64,
	queueDelay int64,
	queueCallback *qlx.QueueCallback,
) *Register {
	queue := qlx.NewQueue(
		cacheAddress,
		identifier,
		queueDelay,
		queueCallback,
	)

	snowprints := SnowprintServiceMap{}
	register := Register{
		routerService: routerService,
		cleanupInterval: cleanupInterval,
		snowprints: snowprints,
		queue: queue,
		cancelCallback: nil,
	}

	return &register
}

func (r *Register) Cancel() {
	r.queue.Cancel()
	if r.cancelCallback != nil {
		(*r.cancelCallback)()
	}
}

// Run this function as a goroutine
func (r *Register) Run() error {
	r.Cancel();
	go r.queue.Run()

	context, cancel := context.WithCancel(context.Background())
	r.cancelCallback = &cancel

	prevNow := time.Now().UnixNano()
	currNow := prevNow
	delta := int64(0)
	
	// loop here
	for {
		select {
		case <-context.Done():
			return context.Err()
		default:
			prevNow = currNow
			currNow = time.Now().UnixNano()

			delta += currNow - prevNow
			if delta > r.cleanupInterval {
				delta %= r.cleanupInterval

				fmt.Println("delta bang!")
				r.CleanupSnowprints(currNow)
				continue
			}
		}
	}

	return nil
}

func (r *Register) UpdateSnowprint(
	details *SnowprintDetails,
) error {
	// if router, add to router
	//
	// if broadcaster
	//    iterate through server broadcasters
	//    if ip updated
	//       update routers
	//
	// if receiver
	//    iterate through server receivers
	//    if ip updated
	//       update each broadcaster service

	return nil
}

func (r *Register) CleanupSnowprints(
	now int64,
) error {
	// iterate through routes
	// /- was routes updated, send it to routes
	//
	// iterate through broadcasters
	// /- was routes updated?
	//    end all broadcasters to routes
	//
	// iterate through routers
	// per-channel
	// /- was channel updated
	// send receiver list to each broadcaster 
	//

	return nil
}

//
// func outside
//
// send all broadcasters to routes
//
// send to broadcaster
//


func createQueueCallback() *qlx.QueueCallback {
	var queueCallback qlx.QueueCallback
	queueCallback = func(
		queuePayload *qlx.QueuePayload,
		cancelCallback *context.CancelFunc,
		err error,
	) error {
		// convert queue payload to snowprint details

		// create request
		// add authorizatin if exists
		// add cookies for loop
		// request body, strings new io
		// timestamp now

		// make request

		return nil
	}

	return &queueCallback
}