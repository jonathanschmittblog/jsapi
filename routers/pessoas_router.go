package routers

import (
	"jsapi/pessoas"
	"jsapi/rabbitmq"

	"github.com/gin-gonic/gin"
)

func CreatePessoa(c *gin.Context) {
	// Cria um novo objeto da pessoa para incluir
	pessoa, err := pessoas.New()
	if err != nil {
		c.JSON(500, 0)
		return 
	}
	// Recebe os novos dados da pessoa do body da requisição
	c.ShouldBind(&pessoa)
	// Efetua a criação da pessoa
	err = pessoa.Create()
	if err != nil {
		c.String(404, err.Error())
		return 
	}
	// Retorna o json da pessoa
	c.JSON(200, pessoa)
	rabbitmq.SendMessage("pessoas", "Cadastro de pessoa: " + pessoa.Nome)
}

func UpdatePessoa(c *gin.Context) {
	// Recebe o nome da pessoa informado no query parameter
	nomePessoa := c.Param("nome")
	if nomePessoa == "" {
		c.String(404, "O nome da pessoa não foi informado nos parâmetros da requisição.")
		return
	}
	// Cria um objeto para a pessoa que será alterada
	pessoa, err := pessoas.New()
	if err != nil {
		c.JSON(500, 0)
		return 
	}
	// Recebe os novos dados da pessoa do body da requisição
	c.ShouldBind(&pessoa)
	// Atualiza os dados da pessoa
	pessoa, err = pessoa.Update(nomePessoa, pessoa)
	if err != nil {
		c.String(404, err.Error())
		return
	}
	// Retorna o json da pessoa
	c.JSON(200, pessoa)
	rabbitmq.SendMessage("pessoas", "Edição de pessoa: " + pessoa.Nome)
}

func GetPessoa(c *gin.Context) {
	// Recebe o nome da pessoa informado no query parameter
	nomePessoa := c.Param("nome")
	if nomePessoa == "" {
		c.String(404, "O nome da pessoa não foi informado nos parâmetros da requisição.")
		return
	}
	// Cria um objeto para a pessoa que será alterada
	pessoa, err := pessoas.New()
	if err != nil {
		c.JSON(500, 0)
		return 
	}
	// Faz a leitura do json da pessoa pelo nome
	err = pessoa.Read(nomePessoa)
	if err != nil {
		c.String(404, "Pessoa '" + nomePessoa + "'não encontrada.")
		return 
	}
	// Retorna o json da pessoa
	c.JSON(200, pessoa)
	rabbitmq.SendMessage("pessoas", "Consulta de pessoa: " + nomePessoa)
}

func DeletePessoa(c *gin.Context) {
	// Recebe o nome da pessoa informado no query parameter
	nomePessoa := c.Param("nome")
	if nomePessoa == "" {
		c.String(404, "O nome da pessoa não foi informado nos parâmetros da requisição.")
		return
	}
	// Cria um objeto para a pessoa que será alterada
	pessoa, err := pessoas.New()
	if err != nil {
		c.JSON(500, 0)
		return 
	}
	// Faz a leitura do json da pessoa pelo nome
	count, err := pessoa.Delete(nomePessoa)
	if err != nil {
		c.String(404, err.Error())
		return 
	}
	// Se não deletou nenhum registro
	if count <= 0 {
		c.String(404, "Pessoa não encontrada para deleção.")
		return 
	}
	// Retorna o json da pessoa
	c.String(200, "Pessoa excluída com sucesso!")
	rabbitmq.SendMessage("pessoas", "Exclusão de pessoa: " + nomePessoa)
}