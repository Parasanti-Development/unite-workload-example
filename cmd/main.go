package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Starting the application...")
	message := flag.String("message", "Hello, This is the default message.", "Message to print")
	flag.Parse()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-ticker.C:
			fmt.Println(*message)
		case sig := <-sigChan:
			fmt.Printf("Received signal to exit: %s\n", sig)
			return
		}
	}
}
