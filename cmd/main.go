package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fsousabt/rate-limiter/internal/notification"
	"github.com/fsousabt/rate-limiter/internal/store"
)

func getCacheAddr() string {
	host := os.Getenv("CACHE_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("CACHE_PORT")
	if port == "" {
		port = "6379"
	}
	return fmt.Sprintf("%s:%s", host, port)
}

func getStore() store.Store {
	storeType := os.Getenv("STORE")
	switch storeType {
	case "memory":
		return store.NewInMemoryStore()
	default:
		return store.NewRedisStore(getCacheAddr())
	}
}

func main() {
	st := getStore()

	ns := notification.NewNotificationServiceImpl(
		notification.NewGateway(),
		notification.NewRateLimiter(st),
	)

	ns.Send(notification.News, "user", "news 1")
	ns.Send(notification.News, "user", "news 2")
	ns.Send(notification.News, "user", "news 3")
	ns.Send(notification.News, "another user", "news 1")

	ns.Send(notification.Status, "user", "status 1")
	ns.Send(notification.Status, "user", "status 2")
	ns.Send(notification.Marketing, "user marketing", "marketing")
	ns.Send(notification.Status, "user", "status 1")

	fmt.Println("Waiting a minute to send another status notification to user")
	time.Sleep(time.Minute)
	ns.Send(notification.Status, "user", "status 4")
}
