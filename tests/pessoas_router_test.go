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
		t.Error("Erro ao instanciar objeto do tipo Pessoa.")
	}
	pessoa.Nome = "Test 123"
	createPessoa(pessoa, t)
	updatePessoa(pessoa.Nome, pessoa, t)
	pessoa2, err := getPessoa(pessoa.Nome, t)
	if err != nil {
		return
	}
	if pessoa2.Nome != pessoa.Nome {
		t.Error("Retorno inesperado na consulta de pessoa.")
	}
	deletePessoa(pessoa.Nome, t)
}

func createPessoa(pessoa *pessoas.Pessoa, t *testing.T) {
	if utils.HttpPost(getHost() + "/pessoas", pessoa) != 200 {
		t.Error("Erro ao criar pessoa.")
	}
}

func updatePessoa(nomePessoa string, pessoa *pessoas.Pessoa, t *testing.T) {
	pessoa.Nome = "Test 321"
	if utils.HttpPut(getHost() + "/pessoas/" + nomePessoa, &pessoa) != 200 {
		t.Error("Erro ao editar pessoa.")
	}
}

func deletePessoa(nome string, t *testing.T) {
	if utils.HttpDelete(getHost() + "/pessoas/" + nome) != 200 {
		t.Error("Erro ao editar pessoa.")
	}
}

func getPessoa(nome string, t *testing.T) (*pessoas.Pessoa, error) {
	status, body := utils.HttpGet(getHost() + "/pessoas/" + nome) 
	if status != 200 {
		t.Error("Erro ao consultar pessoa.")
		return nil, errors.New("Erro ao consultar pessoa.")
	}
	pessoa, err := pessoas.New()
	if err != nil {
		return nil, errors.New("Erro ao criar um objeto de Pessoa.")
	}
	err = json.Unmarshal(body, &pessoa)
	if err != nil {
		return nil, errors.New("Erro ao ler retorno da consulta de pessoa.")
	}
	return pessoa, nil
}

func getHost() string {
	return os.Getenv("JSAPI_HOST") + ":" + os.Getenv("JSAPI_PORT")
}