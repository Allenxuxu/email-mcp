package email

import (
	"crypto/tls"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func Send(receiverEmail, replyTo, cc, bcc []string, senderEmail, subject, text, html, smtpPassword, smtpAddress string) error {
	em := email.NewEmail()
	em.From = senderEmail
	em.To = receiverEmail
	em.Subject = subject
	if len(text) > 0 {
		em.Text = []byte(text)
	}
	if len(html) > 0 {
		em.HTML = []byte(html)
	}
	em.ReplyTo = replyTo
	em.Cc = cc
	em.Bcc = bcc

	return em.SendWithTLS(smtpAddress,
		smtp.PlainAuth("", senderEmail, smtpPassword, ""),
		&tls.Config{InsecureSkipVerify: true})
}
