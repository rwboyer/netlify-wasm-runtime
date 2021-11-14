package util

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

var mailDomain = os.Getenv("MAIL_DOMAIN")
var mailKey = os.Getenv("MAIL_API_KEY")

func Mailer(to []string, from string, sub string, msg []byte, hdrs *map[string]string) error {

	/* Connect to mailgun and make raw message */
	mg := mailgun.NewMailgun(mailDomain, mailKey)
	m := mg.NewMessage(from, sub, string(msg))

	/* Add optional smtp headers */
	for k, v := range *hdrs {
		m.AddHeader(k, v)
	}

	/* Add recipients */
	for _, t := range to {
		err := m.AddRecipient(t)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	/* Have mailgun do it's thing */
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, _, err := mg.Send(ctx, m)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
