package pool_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/narslan/pool"
)

func lambda() {
	fmt.Printf("time: %s\n", time.Now())
}

func TestNewPool(t *testing.T) {
	p := pool.NewPool(4, 10, 1)

	p.Schedule(lambda)
	p.Schedule(lambda)
	time.Sleep(2 * time.Millisecond)

}
