package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const initialWorkersCount = 2

type WorkerPool struct {
	workersCount int
	lastID       int

	mu *sync.Mutex

	input    chan string
	shutdown chan chan struct{}
}

func NewWorkerPool(input chan string) *WorkerPool {
	return &WorkerPool{
		input:    input,
		mu:       &sync.Mutex{},
		shutdown: make(chan chan struct{}),
	}
}

func (pool *WorkerPool) StartWorkers(count int) {
	if pool.shutdown == nil {
		panic("worker pool is shut down")
	}

	pool.mu.Lock()
	defer pool.mu.Unlock()

	for i := 0; i < count; i++ {
		go pool.worker(pool.lastID)
		pool.lastID++
		pool.workersCount++
	}
}

func (pool *WorkerPool) StopWorkers(count int) {
	if pool.shutdown == nil {
		panic("worker pool is shut down")
	}

	pool.mu.Lock()
	defer pool.mu.Unlock()

	if count > pool.workersCount {
		count = pool.workersCount
	}
	shutdownCompleted := make(chan struct{}, count)
	for i := 0; i < count; i++ {
		fmt.Println("sending shutdown signal")
		pool.shutdown <- shutdownCompleted
		pool.workersCount--
	}
	for i := 0; i < count; i++ {
		<-shutdownCompleted
	}
}

func (pool *WorkerPool) Shutdown() {
	pool.StopWorkers(pool.workersCount)
	close(pool.shutdown)
	pool.shutdown = nil
}

func (pool *WorkerPool) worker(id int) {
	fmt.Printf("worker %d is started\n", id)

	for {
		select {
		case completed := <-pool.shutdown:
			fmt.Printf("worker %d is shutted down\n", id)
			completed <- struct{}{}
			return
		case value := <-pool.input:
			fmt.Printf("worker %d: executing work '%s'\n", id, value)
			time.Sleep(5 * time.Second)
			fmt.Printf("worker %d: work '%s' is done\n", id, value)
		}
	}
}

func main() {
	commandChannel := make(chan string)
	pool := NewWorkerPool(commandChannel)
	pool.StartWorkers(initialWorkersCount)

	console := bufio.NewReader(os.Stdin)

	for {
		command, arguments := parseCommand(console)

		switch command {
		case "do":
			commandChannel <- arguments[0]
		case "exit":
			pool.Shutdown()
			close(commandChannel)
			fmt.Println("application completed...")
			return
		case "start":
			pool.StartWorkers(parseCountParameter(arguments))
		case "stop":
			pool.StopWorkers(parseCountParameter(arguments))
		default:
			fmt.Printf("unknown command: %s\n", command)
		}
	}
}

func parseCommand(console *bufio.Reader) (string, []string) {
	input, err := console.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to scan line: %s", err)
	}

	input = strings.TrimRight(input, "\n")
	parameters := strings.Split(input, " ")
	if len(parameters) == 0 {
		return "", nil
	}

	command := parameters[0]
	arguments := parameters[1:]

	return command, arguments
}

func parseCountParameter(arguments []string) int {
	count := 1
	if len(arguments) > 0 {
		count, _ = strconv.Atoi(arguments[0])
	}
	return count
}
