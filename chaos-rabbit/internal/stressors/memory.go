package stressors

import (
	"context"
	"log"
	"time"
)

func ConsumeMemory(ctx context.Context, mb int) {
	if mb <= 0 {
		log.Println("Imvalig memory amount")
		return
	}

	size := mb * 1024 * 1024
	mem := make([]byte, size)

	go func() {
		log.Printf("allocating %dMb of memory\n", mb)
		for {
			select {
			case <-ctx.Done():
				log.Printf("Releasing %dMb of memory\n", mb)
				return
			default:
				for i := 0; i < len(mem); i += 4096 {
					mem[i]++
				}
				time.Sleep(1 * time.Second)
			}
		}
	}()
}
