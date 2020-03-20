FROM golang:alpine3.11 as builder

COPY go.mod go.sum /usr/src/covid0-slackbot/
WORKDIR /usr/src/covid0-slackbot
RUN go mod download

COPY . /usr/src/covid0-slackbot
RUN go build -o /usr/local/bin/covid0-slackbot ./cmd/covid0-slackbot

FROM alpine:3.11
COPY --from=builder /usr/local/bin/covid0-slackbot /usr/local/bin
#RUN mkdir -p /var/covid0-api/temp-storage/covid0.db

CMD [ "/usr/local/bin/covid0-slackbot" ]
