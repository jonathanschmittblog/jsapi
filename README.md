# JSAPI | Modelo de API Go + MongoDB + RabbitMQ + Docker

Olá, seja bem-vindo a minha API Go.
O objetivo desta API é armazenar dados de pessoas em uma collection MongoDB e a cada requisição informar com uma mensagem em uma fila RabbitMQ chamada "pessoas".

A aplicação está configurada para rodar em com Docker.

# Instalação do Docker
- Ubuntu - https://docs.docker.com/install/linux/docker-ce/ubuntu/
- Debian - https://docs.docker.com/install/linux/docker-ce/debian/
- Windows - https://docs.docker.com/docker-for-windows/install/
- Mac - https://docs.docker.com/docker-for-mac/install/

# Configurando a aplicação
As variáveis de ambiente podem ser populadas no arquivo docker-compose.yml nas seções "environment".
- RabbitMQ:
  - RABBITMQ_DEFAULT_USER: "guest"
  - RABBITMQ_DEFAULT_PASS: "guest"
- JSAPI:
  - JSAPI_HOST: "http://localhost"
  - JSAPI_PORT: "3000"
  - JSAPI_RABBITMQ_DIAL: "amqp://guest:guest@rabbitmq:5672/"
  - JSAPI_MONGODB_URI: "mongodb://root:Secret123@mongodb:27017/?retryWrites=true&w=majority"
  - JSAPI_MONGODB_DB_NAME: "jsapidb"
- MongoDB:
  - MONGO_INITDB_ROOT_USERNAME: "root"
  - MONGO_INITDB_ROOT_PASSWORD: "Secret123"
  - MONGO_INITDB_DATABASE: "jsapi"

# Comando para rodar a aplicação com Docker:
- $ docker compose build
- $ docker compose run --rm jsapi go mod init github.com/jonathanschmittblog/jsapi
- $ docker compose run --rm jsapi-ws go mod init github.com/jonathanschmittblog/jsapi-websocket
- $ docker compose run --rm jsapi air init
- $ docker compose up

# Comando para rodar os testes unitários da aplicação com Docker:
- docker exec api go test -run ./...

# Endpoints disponíveis:
- Post: \pessoas
- Get: \pessoas\ :nome_da_pessoa
- Put: \pessoas\ :nome_da_pessoa
- Delete: \pessoas\ :nome_da_pessoa

# Estrutura da Pessoa:
- Obs.: O nome da pessoa não pode repetir
- {"nome": "Viviane", "sexo": "F", "peso": 68.15, "altura": 1.59, "imc": 26,957}
