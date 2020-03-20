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

	if os.Getenv("AWS_ACCESS_KEY_ID") == "" {
		errList = append(errList, errors.New("missing AWS_KEY enviroment"))
	}

	if os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		errList = append(errList, errors.New("missing AWS_SECRET enviroment"))
	}

	if os.Getenv("AWS_S3_REGION") == "" {
		errList = append(errList, errors.New("missing AWS_S3_REGION enviroment"))
	}

	if os.Getenv("AWS_S3_BUCKET") == "" {
		errList = append(errList, errors.New("missing AWS_S3_BUCKET enviroment"))
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

	bot.ExportCSV()
	bot.ExportHTML()
}
