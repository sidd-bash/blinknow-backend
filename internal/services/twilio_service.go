package services

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilioService struct {
	Client      *twilio.RestClient
	ServiceSID  string
}

func NewTwilioService() *TwilioService {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	return &TwilioService{
		Client:     client,
		ServiceSID: os.Getenv("TWILIO_SERVICE_SID"),
	}
}

func (t *TwilioService) SendOTP(phone string) error {
	params := &openapi.CreateVerificationParams{}
	params.SetTo(phone)
	params.SetChannel("sms")
	_, err := t.Client.VerifyV2.CreateVerification(t.ServiceSID, params)
	if err != nil {
		return fmt.Errorf("failed to send OTP: %v", err)
	}
	return nil
}

func (t *TwilioService) VerifyOTP(phone, code string) (bool, error) {
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(phone)
	params.SetCode(code)
	resp, err := t.Client.VerifyV2.CreateVerificationCheck(t.ServiceSID, params)
	if err != nil {
		return false, fmt.Errorf("failed to verify OTP: %v", err)
	}
	return *resp.Status == "approved", nil
}
