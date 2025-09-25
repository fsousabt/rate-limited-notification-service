package store

import "time"

type Rule struct {
	Capacity int64
	Time     time.Duration
}

var RULES_MAP = map[string]Rule{
	"Status":    {Capacity: 2, Time: time.Minute},
	"News":      {Capacity: 1, Time: 24 * time.Hour},
	"Marketing": {Capacity: 3, Time: time.Hour},
}

func GetRulesMap() map[string]Rule {
	return RULES_MAP
}

type Bucket struct {
	Tokens     int64
	LastRefill time.Time
	Rule       Rule
}

type Store interface {
	Get(key string) (Bucket, bool)
	Set(key string, item Bucket) bool
}
