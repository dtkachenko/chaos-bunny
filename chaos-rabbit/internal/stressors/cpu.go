package stressors

import (
	"context"
	"time"
)

// BurnCPU creates load on one CPU core at the given percentage (0–100).
func BurnCPU(ctx context.Context, percent int) {
	if percent <= 0 || percent > 100 {
		return // invalid or no-op
	}

	duration := time.Second // control period
	burn := time.Duration(float64(duration) * float64(percent) / 100.0)
	sleep := duration - burn

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				start := time.Now()
				for time.Since(start) < burn {
					// busy loop (burn CPU)
				}
				time.Sleep(sleep)
			}
		}
	}()
}
