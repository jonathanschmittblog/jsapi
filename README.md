# jsapi
Modelo de API Go + MongoDB + RabbitMQ

Olá, seja bem-vindo a minha API Go.

Requisitos de funcionamento:
- Um banco de dados MongoDB
- Um server RabbitMQ
- Configurar as variáveis de ambiente



Exêmplo de configuração das variáveis de ambiente:

JSAPI_HOST=http://localhost

JSAPI_PORT=3000

JSAPI_RABBITMQ_DIAL=amqp://guest:guest@localhost:5672/

JSAPI_MONGODB_USER=root

JSAPI_MONGODB_PASS=MySecret

JSAPI_MONGODB_DB_NAME=mydbname



Endspoints disponíveis:
- Get
\pessoas\:nome_da_pessoa
- Post
\pessoas
- Put
\pessoas\:nome_da_pessoa
- Delete
\pessoas\:nome_da_pessoa

- Estrutura da Pessoa:
{
    "nome": "",
    "sexo": "",
    "peso": 0,
    "altura": 0,
    "imc": 0
}
