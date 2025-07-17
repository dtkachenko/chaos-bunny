package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dtkachenko/chaos-bunny/chaos-rabbit/internal/stressors"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Stop on Ctrl+C
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go stressors.BurnCPU(ctx, 90)
	go stressors.ConsumeMemory(ctx, 10024)

	log.Println("chaosrabbit running. press Ctrl+C to stop")

	<-sigCh
	log.Println("Signal received. stopping")
	cancel()
}
