package reddo

import (
	"testing"
	"time"
)

func BenchmarkToTime(b *testing.B) {
	now := time.Now()

	for i := 0; i < b.N; i++ {
		_, _ = ToTime(now)

	}
}
