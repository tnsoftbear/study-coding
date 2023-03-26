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
