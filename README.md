# Bot Slack

Repositório criado para interagir com o Slack através de APIs, exportar dados para utilizarmos em comunincações internas e exibir os contribuidores no nosso website.

## users.Export()
Retorna a base de usuários do slack

## users.ExportCSV()
Exporta a base de usuários do slack para `/tmp/covid0-slackbot/output/results.csv`

Vamos usar essa base para comunicação com os membros.

## users.ExportHTML()
Exporta a base de usuários do slack para `/tmp/covid0-slackbot/output/results.html` e envia para o S3 configurado.

Esse arquivo hospedado no S3 é exibido num iFrame dentro do nosso website como forma de agradecimento para todos os que estão participando do projeto.


## Compilando

Tenha a versão 1.13 da SDK GO e rode `make build`

## Empacotando

Rode `make package`, você precisa ter o docker instalado e rodando

## Rodando a aplicação

Copie o arquivo **example.env** e renomei para **.env**, adicionando os valores das variáveis de ambiente.

Rode `make run`
