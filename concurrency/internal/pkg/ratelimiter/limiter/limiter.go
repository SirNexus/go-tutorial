package limiter

import (
	"sync"
	"time"
)

type ThreadSafeLimiter struct {
	sync.RWMutex

	lastPacketTime time.Time
}

func (t *ThreadSafeLimiter) GetLastPacketTime() time.Time {
	return t.lastPacketTime
}

func (t *ThreadSafeLimiter) UpdateLastPacketTime(updated time.Time) {
	t.lastPacketTime = updated
}
