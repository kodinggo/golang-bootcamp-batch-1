package main

import (
	"encoding/json"
	"fmt"
	"log"
	"queue-messaging/task"
	"sync"

	"github.com/hibiken/asynq"
)

func main() {
	// 1. Connect ke redis server
	redisConn := asynq.RedisClientOpt{Addr: "localhost:6379"}

	// 2. Init asynq client (publisher)
	client := asynq.NewClient(redisConn)
	defer client.Close()

	// 5. Enqueue (Publish)
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 20; i++ {
		go func(i int) {
			defer wg.Done()
			// 3. Prepare payload
			payload, _ := json.Marshal(task.SendEmailPayload{
				From:    "john@gmail.com",
				To:      "mark@yahoo.com",
				Subject: fmt.Sprintf("Test Queue %d", i),
				Message: "Mencoba queue messaging",
			})

			// 4. Create task
			queueTask := asynq.NewTask(task.SendEmailTask, payload)

			client.Enqueue(queueTask)
		}(i)
	}

	wg.Wait()

	log.Println("Email berhasil dikirim!")
}
