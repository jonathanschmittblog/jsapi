package main

import (
	"encoding/json"
	"errors"
	"jsapi/pessoas"
	"jsapi/utils"
	"os"
	"testing"
)

func TestPessoa(t *testing.T) {
	pessoa, err := pessoas.New()
	if err != nil {
		t.Error("Erro ao instanciar objeto do tipo Pessoa. Erro: " + err.Error())
	}
	pessoa.Nome = "Test 123"
	pessoa.Altura = 1.70
	pessoa.Peso = 68
	pessoa.Sexo = "M"
	createPessoa(pessoa, t)
	updatePessoa(pessoa.Nome, pessoa, t)
	pessoa2, err := getPessoa(pessoa.Nome, t)
	if err != nil {
		t.Error("Retorno inesperado na consulta de pessoa. Erro:" + err.Error())
	}
	if pessoa2.Nome != pessoa.Nome {
		t.Error("Retorno inesperado na consulta de pessoa.")
	}
	deletePessoa(pessoa.Nome, t)
	_, err = getPessoa(pessoa.Nome, t)
	if err == nil {
		t.Error("Não deveria ter retornado a pessoa na consulta.")
	}
}

func createPessoa(pessoa *pessoas.Pessoa, t *testing.T) {
	err := utils.HttpPost(getHost() + "/pessoas", pessoa)
	if err != nil {
		t.Error("Erro ao criar pessoa. Erro:" + err.Error())
	}
}

func updatePessoa(nomePessoa string, pessoa *pessoas.Pessoa, t *testing.T) {
	pessoa.Nome = "Test 321"
	err := utils.HttpPut(getHost() + "/pessoas/" + nomePessoa, &pessoa)
	if err != nil {
		t.Error("Erro ao editar pessoa. Erro:" + err.Error())
	}
}

func deletePessoa(nome string, t *testing.T) {
	err := utils.HttpDelete(getHost() + "/pessoas/" + nome)
	if err != nil {
		t.Error("Erro ao editar pessoa. Erro:" + err.Error())
	}
}

func getPessoa(nome string, t *testing.T) (*pessoas.Pessoa, error) {
	body, err := utils.HttpGet(getHost() + "/pessoas/" + nome) 
	if err != nil {
		t.Error("Erro ao consultar pessoa. Erro:" + err.Error())
		return nil, errors.New("Erro ao consultar pessoa. Erro:" + err.Error())
	}
	pessoa, err := pessoas.New()
	if err != nil {
		return nil, errors.New("Erro ao criar um objeto de Pessoa." + err.Error())
	}
	err = json.Unmarshal(body, &pessoa)
	if err != nil {
		return nil, errors.New("Erro ao ler retorno da consulta de pessoa." + err.Error())
	}
	return pessoa, nil
}

func getHost() string {
	return os.Getenv("JSAPI_HOST") + ":" + os.Getenv("JSAPI_PORT")
}