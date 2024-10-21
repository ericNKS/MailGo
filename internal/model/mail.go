package model

import "errors"

type Mail struct {
	From     string `json:"from"  binding:"required,max=255"`
	Password string `json:"password"  binding:"required,max=255"`
	SmtpHost string `json:"host"  binding:"required,max=255"`
	SmtpPort int16  `json:"port"  binding:"required"`
}

func CreateMail(from, pass, smtpHost string, smtpPort int16) (*Mail, error) {
	email := &Mail{
		From:     from,
		Password: pass,
		SmtpHost: smtpHost,
		SmtpPort: smtpPort,
	}
	err := validate(email)
	if err != nil {
		return nil, err
	}

	return email, nil
}

func validate(email *Mail) error {

	if email.From == "" {
		return errors.New("from cannot be empty")
	}

	if email.Password == "" {
		return errors.New("password cannot be empty")
	}

	if email.SmtpHost == "" {
		return errors.New("host cannot be empty")
	}

	if email.SmtpPort <= 0 {
		return errors.New("port cannot be less than 0")
	}

	return nil
}
