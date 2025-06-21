package service

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

// SendEmail отправляет электронное письмо с использованием шаблона
func SendCreateEmail(to string, subject string, templatePath string, requestId int32) error {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("Ошибка при парсинге шаблона: %v", err)
		return err
	}
	t.Execute(&body, struct {
		Name      string
		RequestId int32
		OrderLink string
	}{
		Name:      "Кулишко Ярослав",
		RequestId: requestId,
		OrderLink: fmt.Sprintf("https://example.com/orders/%d", requestId),
	})

	err = godotenv.Load()
	if err != nil {
		log.Printf("Ошибка загрузки .env файла: %v", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "yar1k3lfg@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "yar1k3lfg@gmail.com", os.Getenv("GOOGLE_SMTP_PASSWORD"))

	return d.DialAndSend(m)
}
