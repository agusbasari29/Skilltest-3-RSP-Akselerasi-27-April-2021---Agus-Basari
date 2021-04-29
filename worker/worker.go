package worker

import (
	"log"
	"os"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/tasks"
	"github.com/hibiken/asynq"
)

func Workers() {
	r := asynq.RedisClientOpt{Addr: os.Getenv("REDIS_ADDR_PORT")}
	srv := asynq.NewServer(r, asynq.Config{
		Concurrency: 10,
	})

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.PaymentEmail, tasks.HandlePaymentEmailTask)
	mux.HandleFunc(tasks.PaymentPassedEmail, tasks.HandlePaymentPassedEmailTask)
	mux.HandleFunc(tasks.PromotionEmail, tasks.HandlePromotionEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
