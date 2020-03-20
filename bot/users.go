package bot

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"os"
)

type (
	Profile struct {
		RealName string `json:"real_name"`
		Image512 string `json:"image_512"`
		Email    string `json:"email"`
	}

	User struct {
		Name    string  `json:"name"`
		Profile Profile `json:"profile"`
	}

	ExportResponse struct {
		Ok      bool   `json:"ok"`
		Members []User `json:"members"`
	}
)

func Export() bool {

	req, err := http.NewRequest("GET", "https://slack.com/api/users.list", nil)
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg("Error creating request")
		return false
	}

	q := req.URL.Query()
	q.Add("token", os.Getenv("SLACK_TOKEN"))
	q.Add("limit", "1000")

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg(fmt.Sprintf("Error on request: %s", err))
		return false
	}

	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg(fmt.Sprintf("Error on request: %s", err))
		return false
	}

	exportResponse := ExportResponse{}
	err = json.Unmarshal(buf, &exportResponse)
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg(fmt.Sprintf("Error: %s ", err))
		log.Error().Err(err).Str("module", "bot").Msg(fmt.Sprintf("JSON: %s ", buf))
		return false
	}

	var output_path = "/tmp/covid0-slackbot/output"
	err = os.MkdirAll(output_path, os.ModePerm)
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg(fmt.Sprintf("Error creating output directory: %s", err))
		return false
	}

	var output = fmt.Sprintf("%s/result.csv", output_path)
	file, err := os.Create(output)
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg(fmt.Sprintf("Error: %s ", err))
		log.Error().Err(err).Str("module", "bot").Msg(fmt.Sprintf("Error creating file: %s", output))
		return false
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, member := range exportResponse.Members {
		var r []string
		r = append(r, member.Profile.RealName)
		r = append(r, member.Profile.Email)
		r = append(r, member.Profile.Image512)

		err := writer.Write(r)

		if err != nil {
			log.Error().Err(err).Str("module", "bot").Msg(fmt.Sprintf("Error: %s ", err))
			log.Error().Err(err).Str("module", "bot").Msg(fmt.Sprintf("Error writing member: %s", member.Name))
		}
	}

	log.Info().Str("module", "bot").Msg(fmt.Sprintf("Output: %s", output))

	return true
}
