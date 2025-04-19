package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	ServerPort            string
	Dsn                   string
	AppSecret             string
	TwilioAccountSid      string
	TwilioAuthToken       string
	TwilioFromPhoneNumber string
	StripeSecret          string
	PubKey                string
}

func SetupEnv() (cfg AppConfig, err error) {

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("SERVER_PORT")
	Dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	appSecret := os.Getenv("APP_SECRET")

	return AppConfig{ServerPort: httpPort, Dsn: Dsn, AppSecret: appSecret,
		TwilioAccountSid:      os.Getenv("TWILIO_ACCOUNT_SID"),
		TwilioAuthToken:       os.Getenv("TWILIO_AUTH_TOKEN"),
		TwilioFromPhoneNumber: os.Getenv("TWILIO_FROM_PHONE_NUMBER"),
		StripeSecret:          os.Getenv("STRIPE_SECRET"),
		PubKey:                os.Getenv("STRIPE_PUB_KEY"),
	}, nil
}
