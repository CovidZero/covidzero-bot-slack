# Bot Slack

Repositório criado para interagir com o Slack através de APIs

## users.Export()

Exporta a base de usuários do slack para /tmp/covid0-slackbot/output/results.csv
Vamos usar essa base para comunicação com os membros.

## Compilando

Tenha a versão 1.13 da SDK GO e rode `make build`

## Empacotando

Rode `make package`, você precisa ter o docker instalado e rodando

## Rodando a aplicação

Crie um arquivo .env com a variável SLACK_TOKEN.
```.env
SLACK_TOKEN=xxxx-xxxxxxxxx-xxxx
```

Rode `make run`
