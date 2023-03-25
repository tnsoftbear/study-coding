package main

import (
	"context"
	"fmt"

	sendinblue "github.com/sendinblue/APIv3-go-library/v2/lib"
)

func main() {
	var ctx context.Context
	cfg := sendinblue.NewConfiguration()
	//Configure API key authorization: api-key
	apiKey := "<api-key>"
	cfg.AddDefaultHeader("api-key", apiKey)
	//Configure API key authorization: partner-key
	//cfg.AddDefaultHeader("partner-key", apiKey)

	sib := sendinblue.NewAPIClient(cfg)
	sendTransacSms := sendinblue.SendTransacSms{
		Sender:    "I am sender",
		Recipient: "+37122355667",
		Content:   "This is content",
		Type_:     "transactional"}
	result, resp, err := sib.TransactionalSMSApi.SendTransacSms(ctx, sendTransacSms)
	if err != nil {
		fmt.Println("Error when calling TransactionalSMSApi.SendTransacSms: ", err.Error())
		fmt.Println("\nResult: ", result, "\nResponse: ", resp)
		return
	}
	fmt.Println("Result: ", result, "\nResponse: ", resp)
	return
}
