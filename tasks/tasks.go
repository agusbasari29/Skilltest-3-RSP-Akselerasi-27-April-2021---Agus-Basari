package tasks

import (
	"context"
	"strconv"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/hibiken/asynq"
)

const (
	PaymentEmail       = "email:payment"
	PaymentPassedEmail = "email:payment_passed"
	PromotionEmail     = "email:promotion"
)

func NewPaymentEmailTask(email string, amount float32, name string) *asynq.Task {
	payload := map[string]interface{}{"to": email, "amount": amount, "name": name}
	return asynq.NewTask(PaymentEmail, payload)
}

func NewPaymentPassedEmailTask(email string, amount float32, name string) *asynq.Task {
	payload := map[string]interface{}{"to": email, "amount": amount, "name": name}
	return asynq.NewTask(PaymentPassedEmail, payload)
}

func NewPromotionEmailTask(email string, name string) *asynq.Task {
	payload := map[string]interface{}{"to": email, "name": name}
	return asynq.NewTask(PromotionEmail, payload)
}

func HandlePaymentEmailTask(ctx context.Context, t *asynq.Task) error {
	email, _ := t.Payload.GetString("to")
	amount, _ := t.Payload.GetFloat64("amount")
	name, _ := t.Payload.GetString("name")
	subject := "Payment Confirmation"
	body := "Halo " + name + ", Silahkan melakukan pembayaran webinar sebesar " + strconv.Itoa(int(amount))
	helper.SendEmail(email, subject, body)
	return nil
}

func HandlePaymentPassedEmailTask(ctx context.Context, t *asynq.Task) error {
	email, _ := t.Payload.GetString("to")
	amount, _ := t.Payload.GetFloat64("amount")
	name, _ := t.Payload.GetString("name")
	subject := "Payment Comfirmed"
	body := "Halo " + name + ", Terima kasih sudah melakukan pembayaran webinar sebesar " + strconv.Itoa(int(amount))
	helper.SendEmail(email, subject, body)
	return nil
}

func HandlePromotionEmailTask(ctx context.Context, t *asynq.Task) error {
	email, _ := t.Payload.GetString("to")
	name, _ := t.Payload.GetString("name")

	subject := "Payment Confirmation"
	body := "Halo " + name + ", akan ada event menrik nih...!"
	helper.SendEmail(email, subject, body)
	return nil
}
