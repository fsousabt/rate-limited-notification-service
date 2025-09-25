# Rate-Limited Notification Service

A notification service with rate limiting implemented in Go that supports  multiple notification types and different storage backends (Redis or in-memory).

## Features

- Send notifications to users with different types:
  - **Status**: 2 notifications per minute
  - **News**: 1 notification per 24 hours
  - **Marketing**: 3 notifications per hour
- Rate limiting implemented with Token Bucket algorithm
- Supports multiple storage backends: Redis (Default) or In-Memory (for testing)


## Installation and Running

1. Clone the repository:

```bash
git clone https://github.com/fsousabt/rate-limiter.git
cd rate-limiter
```

2. Create a `.env` file in the project root with the following content:

```env
# Choose store backend: "memory" or "redis"
STORE=memory

# Redis configuration (only used if STORE=redis)
CACHE_HOST=redis
CACHE_PORT=6379
```

3. Start the services:

```bash
docker-compose up --build
```

## Running Tests

```bash
go test ./...
```

## Future Improvements

[ ] Add gRPC endpoints

[ ] Deploy service on AWS

[ ] Add metrics dashboard (Prometheus/Grafana)

[ ] Async notifications with queues (Kafka/RabbitMQ)