package stressors

import (
	"context"
	"log"
	"os"
	"time"
)

func DiskWrite(ctx context.Context, path string, iops int) {
	if iops <= 0 {
		log.Println("invalid IOPS value")
		return
	}

	interval := time.Second / time.Duration(iops)
	block := make([]byte, 4096)

	go func() {
		log.Printf("Starting disk write: path=%s iops=%d\n", path, iops)

		for {
			select {
			case <-ctx.Done():
				log.Println("Stopping disk write")
				return
			default:
				start := time.Now()

				f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
				if err != nil {
					log.Printf("disk write error: %v", err)
					time.Sleep(interval)
					continue
				}

				if _, err := f.Write(block); err != nil {
					log.Printf("write failed: %v", err)
				}

				_ = f.Close()

				elapsed := time.Since(start)
				if delay := interval - elapsed; delay > 0 {
					time.Sleep(delay)
				}
			}
		}
	}()
}
