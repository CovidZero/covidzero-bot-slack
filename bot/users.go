package bot

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/zerolog/log"
	"github.com/thoas/go-funk"
	"html/template"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

type (
	Profile struct {
		RealName   string `json:"real_name"`
		StatusText string `json:"status_text"`
		AvatarHash string `json:"avatar_hash"`
		Image192   string `json:"image_192"`
		Image512   string `json:"image_512"`
		Email      string `json:"email"`
	}

	User struct {
		Name    string  `json:"name"`
		Profile Profile `json:"profile"`
		IsBot   bool    `json:"is_bot"`
	}

	ExportResponse struct {
		Ok      bool   `json:"ok"`
		Members []User `json:"members"`
	}
)

func (u User) GetLinkedInLink() string {
	if strings.Contains(u.Profile.StatusText, "https://www.linkedin.com") {
		return u.Profile.StatusText
	}

	return "#"
}

type ByName []User

func (slice ByName) Len() int {
	return len(slice)
}

func (slice ByName) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice ByName) Less(i, j int) bool {
	return slice[i].Profile.RealName < slice[j].Profile.RealName
}

func Export() *ExportResponse {
	req, err := http.NewRequest("GET", "https://slack.com/api/users.list", nil)
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg("Error creating request")
		return nil
	}

	q := req.URL.Query()
	q.Add("token", os.Getenv("SLACK_TOKEN"))
	q.Add("limit", "1000")

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg(fmt.Sprintf("Error on request: %s", err))
		return nil
	}

	var response = ExportResponse{}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&response)
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg("Error decoding JSON")
		return nil
	}

	return &response
}

func ExportCSV() bool {

	var response = Export()
	if response == nil {
		return false
	}

	var outputPath = "/tmp/covid0-slackbot/output"
	err := os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg("Error creating output directory")
		return false
	}

	var output = fmt.Sprintf("%s/result.csv", outputPath)
	file, err := os.Create(output)
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg(fmt.Sprintf("Error creating file: %s", output))
		return false
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, member := range response.Members {
		var r []string
		r = append(r, strconv.FormatBool(member.IsBot))
		r = append(r, member.Profile.RealName)
		r = append(r, member.Profile.StatusText)
		r = append(r, member.Profile.AvatarHash)
		r = append(r, member.Profile.Email)
		r = append(r, member.Profile.Image512)
		r = append(r, member.Profile.Image192)

		err := writer.Write(r)

		if err != nil {
			log.Error().Err(err).Str("module", "bot").Msg(fmt.Sprintf("Error writing member: %s", member.Name))
		}
	}

	log.Info().Str("module", "bot").Msg(fmt.Sprintf("Output: %s", output))

	return true
}

func ExportHTML() bool {
	// Get users from slack
	var response = Export()
	if response == nil {
		return false
	}

	// Create output directory
	var outputPath = "/tmp/covid0-slackbot/output"
	err := os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg("Error creating output directory")
		return false
	}

	// Create output file
	var output = fmt.Sprintf("%s/result.html", outputPath)
	file, err := os.Create(output)
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg(fmt.Sprintf("Error creating file: %s", output))
		return false
	}
	defer file.Close()

	// Get users export template file
	data, err := Asset("data/users_export_template.html")
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg("Error openning template files")
		return false
	}

	// Create new template, ready to render
	t, err := template.New("users_export_template.html").Parse(string(data))
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg("Error parsing template files")
		return false
	}

	// Filter members with photos and non-bots
	var filteredMembers = funk.Filter(response.Members, func(x User) bool {
		if strings.Contains(x.Profile.Image192, "secure.gravatar.com") {
			return false
		}

		if x.IsBot {
			return false
		}

		return true
	})

	// Sort users by name
	sort.Sort(ByName(filteredMembers.([]User)))

	// Execute template and save to output file
	err = t.Execute(file, filteredMembers)
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg("Error executing template")
	}

	log.Info().Str("module", "bot").Msg(fmt.Sprintf("Output: %s", output))

	// Connect to amazon aws services
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_S3_REGION")),
	}))

	// Create a new upload manager to S3
	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		// Define a strategy that will buffer 25 MiB in memory
		u.BufferProvider = s3manager.NewBufferedReadSeekerWriteToPool(25 * 1024 * 1024)
	})

	// Open output file to upload
	file2, err := os.Open(output)
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg("Error opening template file")
	}
	defer file2.Close()

	// Upload file to S3
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:         aws.String("contribuidores.html"),
		Body:        file2,
		ContentType: aws.String("text/html; charset=utf-8"),
	})
	if err != nil {
		log.Error().Err(err).Str("module", "bot").Msg("Error while uploading file to S3")
	}

	return true
}
