package main

import (
	"context"
	"log"
	"queue-messaging/task"

	"github.com/hibiken/asynq"
)

func main() {
	// 1. Connect ke redis server
	redisConn := asynq.RedisClientOpt{Addr: "localhost:6379"}

	// 2. Init asynq server
	asynqServer := asynq.NewServer(redisConn, asynq.Config{
		Concurrency: 10,
	})

	// 3. Init server mux
	mux := asynq.NewServeMux()
	mux.HandleFunc(task.SendEmailTask, func(ctx context.Context, t *asynq.Task) error {
		log.Println(string(t.Payload()))
		return nil
	})

	// 4. Running asynq server
	if err := asynqServer.Run(mux); err != nil {
		log.Fatalf("failed running asynq server, error: %v", err)
	}
}
