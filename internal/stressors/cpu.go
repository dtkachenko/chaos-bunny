package stressors

import (
	"context"
	"log"
	"time"
)

// BurnCPU creates load on one CPU core at the given percentage (0–100).
func BurnCPU(ctx context.Context, percent int) {
	if percent <= 0 || percent > 100 {
		return
	}

	duration := time.Second // control period
	burn := time.Duration(float64(duration) * float64(percent) / 100.0)
	sleep := duration - burn

	go func() {
		log.Printf("Loading cpu to %d Percentage", percent)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				start := time.Now()
				x := 0
				for time.Since(start) < burn {
					x++
					_ = x * x
				}
				time.Sleep(sleep)
			}
		}
	}()
}
