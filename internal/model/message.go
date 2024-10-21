package model

type Message struct {
	Subject string `json:"subject"  binding:"required,max=255"`
	Body    string `json:"body"  binding:"required"`
}

func CreateMessage(subject string, body string) (*Message, error) {
	return &Message{
		Subject: subject,
		Body:    body,
	}, nil
}
