package pessoas

import (
	"context"
	"errors"
	"jsapi/db"
	"jsapi/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Pessoa struct {
	Nome   string `json:"nome"`
	Sexo   string `json:"sexo"`
	Peso   float32 `json:"peso"`
	Altura float32 `json:"altura"`
	Imc    float32 `json:"imc"`
	collection *mongo.Collection
}

// Cria um novo objeto Pessoa
func New() (*Pessoa, error) {
	p := &Pessoa{
		collection: db.GetDatabase().Collection("pessoas"),
	}
	err := p.createIndex()
	return p, err
}

// Cria o índice para definir o Nome da pessoa como único
func (p *Pessoa) createIndex() error {
	_, err := p.collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "nome", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	return utils.HandleError(err)
}

// Adiciona uma pessoa
func (p *Pessoa) Create() error {
	if p.Nome == "" {
		return errors.New("nome da pessoa não informado para edição")
	}
	_, err := p.collection.InsertOne(context.TODO(), p)
	return utils.HandleError(err)
}

// Faz a leitura de uma pessoa a partir do nome
func (p *Pessoa) Read(nome string) error {
	if nome == "" {
		return errors.New("nome da pessoa não informado para consulta")
	}
	filtro := getFilter(nome)
	return p.collection.FindOne(context.TODO(), filtro).Decode(p)
}

// Atualiza os dados de uma pessoa utilizando outra pessoa recebida por parâmetro
func (p *Pessoa) Update(nome string, newPessoa *Pessoa) (*Pessoa, error) {
	if nome == "" || newPessoa.Nome == "" {
		return nil, errors.New("nome da pessoa não informado para edição")
	}
	filtro := getFilter(nome)
	dados := bson.D{primitive.E{Key: "$set", Value: &newPessoa}}
	_, err := p.collection.UpdateOne(context.TODO(), filtro, dados)
	return newPessoa, utils.HandleError(err)
}

// Exclui uma pessoa
func (p *Pessoa) Delete(nome string) (int64, error) {
	if nome == "" {
		return 0, errors.New("nome da pessoa não informado para exclusão")
	}
	filtro := getFilter(nome)
	result, err := p.collection.DeleteOne(context.TODO(), filtro)
	return result.DeletedCount, utils.HandleError(err)
}

// Monta o filtro com o nome da pessoa para utilizar no CRUD
func getFilter(nome string) primitive.D {
	return bson.D{primitive.E{Key: "nome", Value: nome}}
}