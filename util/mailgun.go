package util

import (
	"os"
	"log"
	"time"
	"context"
	"github.com/mailgun/mailgun-go/v4"
)

var mailDomain = os.Getenv("MAIL_DOMAIN")
var mailKey = os.Getenv("MAIL_API_KEY")

func Mailer(to []string, from string, sub string, msg string, hdrs *map[string]string) error {
	mg := mailgun.NewMailgun(mailDomain, mailKey)
	m:= mg.NewMessage(from, sub, msg)
	for _, t := range to {
    err := m.AddRecipient(t)
		if err != nil {
			log.Fatal(err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, _, err := mg.Send(ctx, m)

	if err != nil {  
		log.Fatal(err)
	}

	return nil
}