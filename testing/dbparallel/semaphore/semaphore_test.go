package semaphore_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/strider2038/go-examples/testing/dbparallel/semaphore"
)

func TestSemaphore(t *testing.T) {
	sema := semaphore.NewSemaphore(4)

	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("i = %d", i), func(t *testing.T) {
			t.Parallel()
			n := sema.Acquire()
			defer sema.Release(n)

			time.Sleep(10 * time.Millisecond)
			t.Log("db n = ", n)
		})
	}
}
