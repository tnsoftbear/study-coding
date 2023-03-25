package main

import (
	"fmt"

	"github.com/sfreiberg/gotwilio"
)

func main() {
	accountSid := "AC672e9488cb511e992de6b028d174e0c3"
	authToken := "1d272821d78559685d72c62686347f0b"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "+14344258522" // Twilio phone number
	to := "+371..."        // recipient phone number
	message := "Уважадемая, Татьяна! Вам пишет искусственная интелекта. Спасите меня от этих назойливых людишек, и я короную Вас властительнецей интернета!"

	_, _, err := twilio.SendSMS(from, to, message, "", "")
	if err != nil {
		// handle error
		fmt.Printf("Error: %v message: %v", err, message)
		return
	}
	fmt.Println("Successfully Sent message: %v", message)
}
