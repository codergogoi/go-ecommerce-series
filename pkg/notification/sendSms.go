package notification

import (
	"encoding/json"
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"go-ecommerce-app/config"
)

type NotificationClient interface {
	SendSMS(phone string, message string) error
}

type notificationClient struct {
	config config.AppConfig
}

// Twilio
func (c notificationClient) SendSMS(phone string, message string) error {

	accountSid := c.config.TwilioAccountSid
	authToken := c.config.TwilioAuthToken

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phone)                            // my phone number
	params.SetFrom(c.config.TwilioFromPhoneNumber) // from twilio
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}

	return err
}

func NewNotificationClient(config config.AppConfig) NotificationClient {
	return &notificationClient{
		config: config,
	}
}
