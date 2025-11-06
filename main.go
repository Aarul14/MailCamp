package main

import (
	"sync"
)
type Recipient struct {
	Name string
	Email string
}
func main(){
	recipientChannel := make(chan Recipient)
    var wg sync.WaitGroup

	go loadRecipient("./email.csv", recipientChannel)

	var worker = 5 //fomring 5 worker as consumer
	for   i:=0 ; i<worker; i++{
		wg.Add(1) //for each worker of consumer
		go emailWorker(i, recipientChannel, &wg)
	}

	wg.Wait()

}