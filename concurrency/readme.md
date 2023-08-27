# Rate Limiter

Rate Limiter is a demonstration of golang using channels, waitgroups, and shared memory. The rate limiter algorithm itself is not complicated, but this example shows how multiple threads can work in tandem to accomplish some work based on shared state.

```
2023/08/27 11:40:43 worker 1 event rate limited!

2023/08/27 11:40:44 worker 0 event passed successfully

2023/08/27 11:40:44 worker 1 event rate limited!

2023/08/27 11:40:44 worker 0 event rate limited!

2023/08/27 11:40:45 worker 1 event rate limited!
```

Each worker must maintain an accurate state of the world on when the last successful request was allowed through so that it can accurately maintain the global rate limit.

Run with: 

```
go run cmd/rate-limiter/main.go
```
