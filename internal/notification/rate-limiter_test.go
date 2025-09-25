package notification

import (
	"testing"

	"github.com/fsousabt/rate-limiter/internal/store"
)

func TestRateLimiter_Allow_AllNotificationTypes(t *testing.T) {
	memStore := store.NewInMemoryStore()
	rl := NewRateLimiter(memStore)

	userId := "user1"

	notificationTypes := []NotificationType{Status, News, Marketing}
	for _, notType := range notificationTypes {
		rule := rl.rules[notType.String()]

		for i := int64(0); i < rule.Capacity; i++ {
			allowed := rl.Allow(userId, notType)
			if !allowed {
				t.Fatalf("[%s] expected Allow=true, but got false on attempt %d", notType.String(), i+1)
			}
		}

		allowed := rl.Allow(userId, notType)
		if allowed {
			t.Fatalf("[%s] expected Allow=false, but got true", notType.String())
		}

		bucketKey := rl.makeKey(userId, notType)
		bucket, _ := memStore.Get(bucketKey)
		bucket.LastRefill = bucket.LastRefill.Add(-rule.Time)
		memStore.Set(bucketKey, bucket)

		allowed = rl.Allow(userId, notType)
		if !allowed {
			t.Fatalf("[%s] expected Allow=true after refill, but got false", notType.String())
		}
	}
}
