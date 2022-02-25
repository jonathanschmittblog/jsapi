# JSAPI | Modelo de API Go + MongoDB + RabbitMQ + Docker

Olá, seja bem-vindo a minha API Go.
O objetivo desta API é armazenar dados de pessoas em uma collection MongoDB e a cada requisição informar com uma mensagem em uma fila RabbitMQ chamada "pessoas".

Esta aplicação utiliza do outro projeto (jsapi-websocket) para exibir as mensagens do RabbitMQ no navegador.
- https://github.com/jonathanschmittblog/jsapi

A aplicação está configurada para rodar em com Docker.

# Instalação do Docker
- Ubuntu - https://docs.docker.com/install/linux/docker-ce/ubuntu/
- Debian - https://docs.docker.com/install/linux/docker-ce/debian/
- Windows - https://docs.docker.com/docker-for-windows/install/
- Mac - https://docs.docker.com/docker-for-mac/install/

# Configurando a aplicação
As variáveis de ambiente podem ser populadas no arquivo docker-compose.yml nas seções "environment".
- RabbitMQ:
  - RABBITMQ_DEFAULT_USER: "jsuser"
  - RABBITMQ_DEFAULT_PASS: "jspass"
- JSAPI:
  - JSAPI_HOST: "http://localhost"
  - JSAPI_PORT: "3000"
  - JSAPI_RABBITMQ_DIAL: "amqp://jsuser:jspass@rabbitmq:5672/"
  - JSAPI_MONGODB_URI: "mongodb://root:Secret123@mongodb:27017/?retryWrites=true&w=majority"
  - JSAPI_MONGODB_DB_NAME: "jsapidb"
- JSAPI-WEBSOCKET:
  - JSAPIWS_PORT: "3001"
  - JSAPIWS_RABBITMQ_DIAL: "amqp://jsuser:jspass@rabbitmq:5672/"
- MongoDB:
  - MONGO_INITDB_ROOT_USERNAME: "root"
  - MONGO_INITDB_ROOT_PASSWORD: "Secret123"
  - MONGO_INITDB_DATABASE: "jsapi"

# Comando para rodar a aplicação com Docker:
- $ docker compose up --no-cache

# Comando para rodar os testes unitários da aplicação:
- $ docker exec jsapi-1 go test -v ./test

# Endpoints disponíveis:
- Post: \pessoas
- Get: \pessoas\ :nome_da_pessoa
- Put: \pessoas\ :nome_da_pessoa
- Delete: \pessoas\ :nome_da_pessoa

# Estrutura da Pessoa:
- Obs.: O nome da pessoa não pode repetir
- {"nome": "Viviane", "sexo": "F", "peso": 68.15, "altura": 1.59, "imc": 26,957}
