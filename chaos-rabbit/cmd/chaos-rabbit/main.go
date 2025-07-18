package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dtkachenko/chaos-bunny/chaos-rabbit/internal/stressors"
)

func main() {
	cpuPercent := flag.Int("cpu", 0, "CPU usage percent (0-100)")
	memMb := flag.Int("mem", 0, "Amount of memory to consume in MB")
	discIOPS := flag.Int("iops", 0, "Disk IOPS")
	discPath := flag.String("path", "/tmp/chaos-rabbit", "Disk path to write into")

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Stop on Ctrl+C
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	if *cpuPercent > 0 {
		go stressors.BurnCPU(ctx, *cpuPercent)
	}
	if *memMb > 0 {
		go stressors.ConsumeMemory(ctx, *memMb)
	}
	if *discIOPS > 0 {
		go stressors.DiskWrite(ctx, *discPath, *discIOPS)
	}

	log.Println("chaosrabbit running. press Ctrl+C to stop")

	<-sigCh
	log.Println("Signal received. stopping")
	cancel()
}
