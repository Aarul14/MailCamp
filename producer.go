package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func loadRecipient(filePath string, ch chan Recipient) error {

	f, err := os.Open(filePath)
	if err != nil{
		return err
	}

	defer f.Close()

	r := csv.NewReader(f)
	records, err  := r.ReadAll()  
	if err != nil {
		 return err
	}


	for _, record:= range records[1:]{  //ignored the first row of csv that is the header row
		fmt.Println(record)
		//send -> consumer -> channels

		//untill and unless there is no one to recieve
		//this remains blocked

		ch <-  Recipient{   
			Name: record[0],
			Email: record[1],
		}
		
	}


	return nil
}