package registryxs

import (
	// "os"
	"context"
	"fmt"
	"testing"
	"time"

	qlx "github.com/herebythere/queuelx/v0.1/golang"
)

const (
	testIdentifier   = "test_registry"
	testRouter = "test_router"
	cleanupInterval = int64(2 * time.Second)
	queueDelay = int64(2 * time.Second)
)

var (
	// localCacheAddress = os.Getenv("LOCAL_CACHE_ADDRESS")
	localCacheAddress = "http://10.88.0.1:6050"
)

func TestCreateRegister(t *testing.T) {
	var testQueueCallback qlx.QueueCallback
	testQueueCallback = func (
		payload *qlx.QueuePayload,
		cancelCallback *context.CancelFunc,
		err error,
	) error {
		fmt.Println(payload)
		return nil
	}

	registry := NewRegistry(
		localCacheAddress,
		testIdentifier,
		testRouter,
		cleanupInterval,
		queueDelay,
		&testQueueCallback,
	)

	if registry == nil {
		t.Fail()
		t.Logf("registry was nil")
	}
}