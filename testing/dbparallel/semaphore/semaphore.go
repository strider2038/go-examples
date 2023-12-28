package semaphore

import (
	"os"
	"strconv"
	"sync"
)

func Acquire() int {
	return semaphore.Acquire()
}

var semaphore = sync.OnceValue(func() *Semaphore {
	count, _ := strconv.Atoi(os.Getenv("DATABASE_COUNT"))
	if count < 1 || count > 32 {
		count = 1
	}

	return NewSemaphore(count)
})()

func Release(n int) {
	semaphore.Release(n)
}

type Semaphore struct {
	freeDB chan int
}

func NewSemaphore(count int) *Semaphore {
	s := &Semaphore{freeDB: make(chan int, count)}
	for i := 0; i < count; i++ {
		s.freeDB <- i
	}

	return s
}

func (s *Semaphore) Acquire() int {
	n := <-s.freeDB

	return n
}

func (s *Semaphore) Release(n int) {
	s.freeDB <- n
}
