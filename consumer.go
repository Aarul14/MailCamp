package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {

	defer wg.Done() //when sare email read hogye from the channel, close the channel

	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"
		// formattedMessage := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n", recipient.Email,"This is a test email")
		// msg := []byte(formattedMessage)

		msg, err := executeTemplate(recipient)
		if err != nil {
			fmt.Printf("worker %d: Error parsing template to %s", id, recipient.Email)
			//todo: add to dlq
			continue
		}

		from := "aarul@mailcamp.com"
		to := recipient.Email
		subject := "Welcome to MailCamp 🎉"

		// ✅ Construct a valid MIME message with headers
		email := fmt.Sprintf(
			"From: %s\r\n"+
				"To: %s\r\n"+
				"Subject: %s\r\n"+
				"MIME-Version: 1.0\r\n"+
				"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
				"\r\n%s",
			from, to, subject, msg,
		)

		fmt.Printf("%d) Sending email to %s having name as %s \n", id, recipient.Email, recipient.Name)

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "aarul@mailcamp.com", []string{recipient.Email}, []byte(email))
		if err != nil {
			log.Printf("⚠️ Worker %d: failed to send to %s: %v", id, recipient.Email, err)
		}

		time.Sleep(18 * time.Millisecond) //diff btw time sending another mail

		fmt.Printf("%d) Sent email to %s having name as %s \n", id, recipient.Email, recipient.Name)

	}

}
