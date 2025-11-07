package main

import (
	"bytes"
	"sync"
	"text/template"
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


func executeTemplate(r Recipient)(string, error){
	t, err := template.ParseFiles("email.tmpl")
	if err != nil{
		return "", err
	} 
	
	//buffer is used to store the htmlTemplate in a memory

	var tpl bytes.Buffer
	
	//Take the template t, fill in all {{.Name}} and other variables using r, and write the result into tpl.”
	err = t.Execute(&tpl, r) 
	if err!=nil{
		return "", err
	}

	//tpl.String() converts the buffer into a plain Go string.
	return tpl.String(),nil

}