package main

import (
	"encoding/json"
	"log"
	"sendMail/internal/model"
	"sendMail/internal/service"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	con, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panicf("%s: %s", "Falha ao fazer a conexao", err)
	}
	defer con.Close()

	ch, err := con.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Falha ao abrir um channel", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"sendMail",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panicf("%s: %s", "Falha ao declarar a queue", err)
	}

	msg, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panicf("%s: %s", "Error", err)
	}

	var forever chan struct{}

	go func() {
		for d := range msg {
			var cp model.Campaign

			err := json.Unmarshal(d.Body, &cp)
			if err != nil {
				log.Printf("Error: %s", err.Error())
				continue
			}

			go func() {
				for _, d := range cp.Destinatarios {
					service.ExecuteSendMail(&cp.Remetente, &cp.Mensagem, d)
				}
			}()

			log.Printf("Success: %+v", cp.Remetente)
		}
	}()
	<-forever
}
