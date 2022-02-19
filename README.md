# jsapi
- Modelo de API Go + MongoDB + RabbitMQ + Docker

Olá, seja bem-vindo a minha API Go.

A aplicação está configurada para rodar em com Docker.

Instalação do Docker
- Ubuntu - https://docs.docker.com/install/linux/docker-ce/ubuntu/
- Debian - https://docs.docker.com/install/linux/docker-ce/debian/
- Windows - https://docs.docker.com/docker-for-windows/install/
- Mac - https://docs.docker.com/docker-for-mac/install/

Depois de instalar o Docker, utilize os comandos:
- $ docker compose build
- $ docker compose run --rm jsapi go mod init github.com/jonathanschmittblog/jsapi
- $ docker compose run --rm app air init
- $ docker compose up

Endspoints disponíveis:
- Post: \pessoas
- Get: \pessoas\:nome_da_pessoa
- Put: \pessoas\:nome_da_pessoa
- Delete: \pessoas\:nome_da_pessoa

Estrutura da Pessoa:
- {"nome": "", "sexo": "", "peso": 0, "altura": 0, "imc": 0}
