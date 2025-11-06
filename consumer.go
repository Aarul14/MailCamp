package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)


func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup){

	defer wg.Done()   //when sare email read hogye from the channel, close the channel



	for recipient := range ch {
			smtpHost := "localhost"
			smtpPort := "1025"
			formattedMessage := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n", recipient.Email,"This is a test email")
			msg := []byte(formattedMessage)

			fmt.Printf("%d) Sending email to %s having name as %s \n", id,recipient.Email, recipient.Name)

			err:= smtp.SendMail(smtpHost+":"+smtpPort, nil, "aarul@mailcamp.com", []string{recipient.Email}, msg )
			if err != nil{
				log.Printf("⚠️ Worker %d: failed to send to %s: %v", id, recipient.Email, err)
			}

			time.Sleep(18 * time.Millisecond) //diff btw time sending another mail

			fmt.Printf("%d) Sent email to %s having name as %s \n", id,recipient.Email, recipient.Name)


	}

	
}

