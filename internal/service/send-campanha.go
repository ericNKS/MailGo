package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"sendMail/internal/model"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type createCampaign struct {
	campaign *model.Campaign
}

func validateCampaign(email *model.Mail, to *[]string, mensagem *model.Message) error {
	if email == nil {
		return errors.New("email is required")
	}

	if len(*to) < 1 {
		return errors.New("need someone to received")
	}

	if mensagem == nil {
		return errors.New("body menssage cannot be empty")
	}

	return nil
}

func CreateCampaign(email *model.Mail, to []string, mensagem *model.Message) (*createCampaign, error) {
	err := validateCampaign(email, &to, mensagem)
	if err != nil {
		return nil, err
	}

	campaign := &createCampaign{
		campaign: &model.Campaign{
			Remetente:     *email,
			Destinatarios: to,
			Mensagem:      *mensagem,
		},
	}

	return campaign, nil
}

func (campaign *createCampaign) Execute(wg *sync.WaitGroup) error {
	defer wg.Done()
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a channel", err)
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"sendMail", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to declare a queue", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(campaign.campaign)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Panicf("%s: %s", "Failed to publish a message", err)
		return err
	}

	return nil
}
