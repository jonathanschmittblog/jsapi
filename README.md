# JSAPI | Modelo de API Go + MongoDB + RabbitMQ + Docker

Olá, seja bem-vindo a minha API Go.

A aplicação está configurada para rodar em com Docker.

# Instalação do Docker
- Ubuntu - https://docs.docker.com/install/linux/docker-ce/ubuntu/
- Debian - https://docs.docker.com/install/linux/docker-ce/debian/
- Windows - https://docs.docker.com/docker-for-windows/install/
- Mac - https://docs.docker.com/docker-for-mac/install/

# Comando para rodar a aplicação com Docker:
- $ docker compose build
- $ docker compose run --rm api go mod init github.com/jonathanschmittblog/jsapi
- $ docker compose run --rm api air init
- $ docker compose up

# Comando para rodar os testes unitários da aplicação com Docker:
- docker exec api go test -run ./...

# Endpoints disponíveis:
- Post: \pessoas
- Get: \pessoas\ :nome_da_pessoa
- Put: \pessoas\ :nome_da_pessoa
- Delete: \pessoas\ :nome_da_pessoa

# Estrutura da Pessoa:
{"nome": "Viviane", "sexo": "F", "peso": 68.15, "altura": 1.59, "imc": 26,957}
