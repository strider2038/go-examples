package pkg1

import (
	"testing"
	"time"

	"github.com/strider2038/go-examples/testing/dbparallel/semaphore"
)

func TestP(t *testing.T) {
	n := semaphore.Acquire()
	defer semaphore.Release(n)

	t.Log("db n = ", n, time.Now().String())
	time.Sleep(time.Second)
}
