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
	cfg.AddDefaultHeader("partner-key", apiKey)

	sib := sendinblue.NewAPIClient(cfg)
	result, resp, err := sib.AccountApi.GetAccount(ctx)
	if err != nil {
		fmt.Println("Error when calling AccountApi->get_account: ", err.Error())
		return
	}
	fmt.Println("GetAccount Object:", result, " GetAccount Response: ", resp)
	return
}
