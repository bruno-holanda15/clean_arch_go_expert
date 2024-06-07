package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"clean_arch_desafio/pkg/events"

	"github.com/streadway/amqp"
)

type OrdersListedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrdersListedHandler(rabbitMQChannel *amqp.Channel) *OrdersListedHandler {
	return &OrdersListedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrdersListedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("All orders listed: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
