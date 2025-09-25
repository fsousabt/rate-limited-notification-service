package notification

import (
	"fmt"
	"time"

	"github.com/fsousabt/rate-limiter/internal/store"
)

type RateLimiter struct {
	store store.Store
	rules map[string]store.Rule
}

func NewRateLimiter(s store.Store) *RateLimiter {
	return &RateLimiter{store: s, rules: store.GetRulesMap()}
}

func (r RateLimiter) makeKey(userId string, notType NotificationType) string {
	return fmt.Sprintf("rl:%s:%s", userId, notType.String())
}

func (r *RateLimiter) Allow(userId string, notType NotificationType) bool {
	rule, foundRule := r.rules[notType.String()]
	if !foundRule {
		return false
	}

	key := r.makeKey(userId, notType)
	bucket, ok := r.store.Get(key)

	if !ok {
		bucket = store.Bucket{
			Tokens:     rule.Capacity,
			LastRefill: time.Now(),
			Rule: store.Rule{
				Capacity: rule.Capacity,
				Time:     rule.Time,
			},
		}
	}

	now := time.Now()
	elapsed := now.Sub(bucket.LastRefill)
	refillTokens := elapsed.Nanoseconds() * bucket.Rule.Capacity / bucket.Rule.Time.Nanoseconds()
	if refillTokens > 0 {
		bucket.Tokens = min(bucket.Rule.Capacity, bucket.Tokens+refillTokens)
		bucket.LastRefill = now
	}

	if bucket.Tokens > 0 {
		bucket.Tokens--
		r.store.Set(key, bucket)
		return true
	}

	r.store.Set(key, bucket)
	return false

}
