package main

import (
	"encoding/json"
	"log"
	"queue-messaging/task"

	"github.com/hibiken/asynq"
)

func main() {
	// 1. Connect ke redis server
	redisConn := asynq.RedisClientOpt{Addr: "localhost:6379"}

	// 2. Init asynq client (publisher)
	client := asynq.NewClient(redisConn)
	defer client.Close()

	// 3. Prepare payload
	payload, _ := json.Marshal(task.SendEmailPayload{
		From:    "john@gmail.com",
		To:      "mark@yahoo.com",
		Subject: "Test Queue",
		Message: "Mencoba queue messaging",
	})

	// 4. Create task
	queueTask := asynq.NewTask(task.SendEmailTask, payload)

	// 5. Enqueue (Publish)
	client.Enqueue(queueTask)

	log.Println("Email berhasil dikirim!")
}
