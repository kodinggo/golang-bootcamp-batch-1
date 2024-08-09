package main

import (
	"context"
	"encoding/json"
	"log"
	"queue-messaging/task"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	// 1. Connect ke redis server
	redisConn := asynq.RedisClientOpt{Addr: "localhost:6379"}

	// 2. Init asynq server
	asynqServer := asynq.NewServer(redisConn, asynq.Config{
		Concurrency: 2,
	})

	// 3. Init server mux
	mux := asynq.NewServeMux()
	mux.HandleFunc(task.SendEmailTask, func(ctx context.Context, t *asynq.Task) error {
		time.Sleep(2 * time.Second)

		var payload task.SendEmailPayload
		json.Unmarshal(t.Payload(), &payload)

		log.Printf("Handle send email for subject %s", payload.Subject)

		return nil
	})

	// 4. Running asynq server
	if err := asynqServer.Run(mux); err != nil {
		log.Fatalf("failed running asynq server, error: %v", err)
	}
}
