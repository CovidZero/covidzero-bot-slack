package main

import (
	"errors"
	"flag"
	"github.com/CovidZero/covid0-slackbot/bot"
	"github.com/rs/zerolog/log"
	"os"
)

func SaneEnv() []error {
	var errList []error

	if os.Getenv("SLACK_TOKEN") == "" {
		errList = append(errList, errors.New("missing SLACK_TOKEN enviroment"))
	}

	return errList
}

func main() {
	flag.Parse()

	if errList := SaneEnv(); errList != nil {
		for _, e := range errList {
			log.Error().Err(e).Msg("Check if your environment configuration is correct")
		}
		log.Fatal().Msg("Your environment is not properly configured. Please check the log for more information")
	}

	bot.Export()
}
