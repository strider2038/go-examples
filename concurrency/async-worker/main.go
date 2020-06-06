package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func worker(shutdownChannel <-chan struct{}, commandChannel <-chan string) {
	for {
		select {
		// completes work on shutdown signal
		case <-shutdownChannel:
			fmt.Println("worker: shutdown signal received")
			return
		// executes command
		case command := <-commandChannel:
			result := handleCommand(command)
			fmt.Println("worker:", result)
		}
	}
}

func handleCommand(command string) string {
	arguments := strings.Split(command, " ")
	if len(arguments) <= 0 {
		return "empty command"
	}

	switch arguments[0] {
	case "countdown":
		return handleCountdown(arguments)
	default:
		return fmt.Sprintf("unknown command: %s", arguments[0])
	}
}

func handleCountdown(arguments []string) string {
	var err error

	countdownDuration := 1
	if len(arguments) >= 2 {
		countdownDuration, err = strconv.Atoi(arguments[1])
		if err != nil {
			return fmt.Sprintf("invalid argument for 'countdown' command: %s", err)
		}
	}

	fmt.Printf("worker: starting countdown for %d seconds\n", countdownDuration)

	for i := countdownDuration; i >= 0; i-- {
		time.Sleep(time.Second)
		fmt.Printf("worker: %d seconds left\n", i)
	}

	return "countdown completed"
}

func main() {
	shutdownChannel := make(chan struct{})
	commandChannel := make(chan string)

	go worker(shutdownChannel, commandChannel)

	reader := bufio.NewReader(os.Stdin)

	for {
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("failed to scan line: %s", err)
		}

		command = strings.TrimRight(command, "\n")

		if command == "exit" {
			fmt.Println("sending shutdown signal to worker...")
			shutdownChannel <- struct{}{}
			fmt.Println("shutdown completed...")

			break
		}

		commandChannel <- command
	}
}
