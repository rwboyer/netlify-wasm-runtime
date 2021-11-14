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

type TxtMailer struct {
	mg *mailgun.MailgunImpl
	m  *mailgun.Message
}

type HtmlMailer struct {
	*TxtMailer
}

type Mailer interface {
	Send(to []string, from string, sub string, msg string, hdrs *map[string]string) error
}

func NewTextMailer(to []string, from string, sub string, msg string, hdrs *map[string]string) (*TxtMailer, error) {

	var t TxtMailer

	/* Connect to mailgun and make raw message */
	t.mg = mailgun.NewMailgun(mailDomain, mailKey)
	t.m = t.mg.NewMessage(from, sub, string(msg))

	/* Add optional smtp headers */
	for k, v := range *hdrs {
		t.m.AddHeader(k, v)
	}

	/* Add recipients */
	for _, r := range to {
		err := t.m.AddRecipient(r)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return &t, nil
}

func NewHtmlMailer(to []string, from string, sub string, msg string, hdrs *map[string]string) (*HtmlMailer, error) {
	var hm HtmlMailer
	var err error

	if hm.TxtMailer, err = NewTextMailer(to, from, sub, msg, hdrs); err != nil {
		log.Println(err)
		return nil, err
	}

	return &hm, nil
}

func (hm *HtmlMailer) Send(msg string) error {
	hm.m.SetHtml(msg)
	hm.TxtMailer.Send(msg)
	return nil
}

func (tm *TxtMailer) Send(msg string) error {

	/* Have mailgun do it's thing */
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, _, err := tm.mg.Send(ctx, tm.m)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
