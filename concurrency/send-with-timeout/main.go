package main

import (
	"log"
	"strconv"
	"sync"
	"time"
)

func main() {
	// отправка в канал с таймаутом (эмуляция переполнения очереди)
	log.Println("start")

	ch := make(chan string, 1)
	w := sync.WaitGroup{}
	w.Add(1)

	go func() {
		for s := range ch {
			log.Println("receive", s)
			time.Sleep(2 * time.Second)
			log.Println("complete", s)
		}
		w.Done()
	}()

	for i := 0; i < 3; i++ {
		select {
		case ch <- strconv.Itoa(i):
			log.Println("send", strconv.Itoa(i))
		case <-time.After(time.Second):
			log.Println("timeout", strconv.Itoa(i))
		}
	}

	close(ch)
	w.Wait()

	log.Println("finish")
}
