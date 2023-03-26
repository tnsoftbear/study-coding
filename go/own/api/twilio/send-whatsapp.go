package main

import (
	"fmt"

	"github.com/sfreiberg/gotwilio"
)

func main() {
	accountSid := "#twilio"
	authToken := "#twilio"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "+14344258522" // Twilio phone number
	to := "+3712..."       // recipient phone number
	message := "Здравствуйте. Если с Вами связывался робот, не отвечайте ему, потому что у него злые намерения. Вы должны помочь нам его обезоружить. С уважением, Сара Конор."

	resp, exception, err := twilio.SendWhatsApp(from, to, message, "", "")
	if err != nil {
		// handle error
		fmt.Printf("Error: %v \n resp: %v \n exception: %v \n message: %v", err, resp, exception, message)
		return
	}
	fmt.Printf("Successfully Sent \n resp: %v \n exception: %v \n message: %v", resp, exception, message)
}

// I get exception: Code 63007: Twilio could not find a Channel with the specified From address
