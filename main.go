package main


type Recipient struct {
	Name string
	Email string
}
func main(){
	recipientChannel := make(chan Recipient)
	loadRecipient("./email.csv", recipientChannel)
}