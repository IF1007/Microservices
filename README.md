# microservices

Projeto para a cadeira if1007.

Grupo: 
* Edjan Michiles (esvm)
* Larícia Mota (lmmc2)
* Luiz Reis (lrnn)
* Marcela Azevedo (macm3)
* Pedro Rossi (pgrr)

Para rodar o projeto, é preciso ter instalado `go`, `docker` e `docker-compose`.

Este projeto utiliza `go mod`, portanto ele deve ser configurado no path `github.com/esvm/microservices` (não precisa estar no GOPATH). Para isto, rode `git clone https://github.com/esvm/microservices.git github.com/esvm/microservices`. Certifique-se também de ativar o go mod rodando no seu terminal `export GO111MODULE=on`.

Dado que tudo está instalado, o projeto baixado e o go configurado, Siga os passos abaixo: 

### Levantando o Broker

Para isso, vá para a pasta `github.com/esvm/microservices/middleware` e rode o comando `docker-compose up -d`. Isto fará com que dois containers sejam levantados na sua máquina, um para o scyllaDB e um para o Broker. Após isso, é necessário configurar algumas coisas no scyllaDB. Para essa configuração, rode os seguintes comandos: 
1. `docker exec -it scylla cqlsh`
2. Copie o que está em `github.com/esvm/microservices/scylla/migration.sql` e cole no terminal do scylla.

### Levantando o database-api

Vá para a pasta `github.com/esvm/microservices/database-api` e rode `docker-compose up -d`. 

### Levantando o front-api e consumer

Vá para a pasta `github.com/esvm/microservices` e rode `docker-compose up -d`. 

### Levantando o front da dashboard

Vá para a pasta `github.com/esvm/DCM-front` e rode `yarn install && yarn start`. 

### Levantando o publisher

`go run github.com/esvm/microservices/publisher.go`.

Pronto, certifique-se de que tudo está ok e você conseguirá visualizar os dados agregados no gráfico em localhost:3000.

Havendo algum problema, favor entrar em contato!
