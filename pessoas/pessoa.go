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
	Peso   float64 `json:"peso"`
	Altura float64 `json:"altura"`
	Imc    float64 `json:"imc"`
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
	return err
}

func (p *Pessoa) ValidaPessoa() error {
	if p.Nome == "" {
		return errors.New("o nome da pessoa não informado")
	} else if p.Peso == 0 {
		return errors.New("o peso da pessoa não informado")
	} else if p.Altura == 0 {
		return errors.New("a altura da pessoa não informado")
	} else if p.Sexo != "F" && p.Sexo != "M"  {
		return errors.New("o sexo da pessoa não informado corretamente [F/M]")
	}
	return nil
}

// Adiciona uma pessoa
func (p *Pessoa) Create() error {
	errVal := p.ValidaPessoa()
	if errVal != nil {
		return errVal
	}
	p.Imc = utils.CalculaImc(p.Peso, p.Altura)
	_, err := p.collection.InsertOne(context.TODO(), p)
	return err
}

// Faz a leitura de uma pessoa a partir do nome
func (p *Pessoa) Read(nome string) error {
	if nome == "" {
		return errors.New("nome da pessoa não informado para edição")
	}
	filtro := getFilter(nome)
	return p.collection.FindOne(context.TODO(), filtro).Decode(p)
}

// Atualiza os dados de uma pessoa utilizando outra pessoa recebida por parâmetro
func (p *Pessoa) Update(nome string, newPessoa *Pessoa) (*Pessoa, error) {
	if nome == "" {
		return nil, errors.New("nome da pessoa não informado para edição")
	}
	errVal := p.ValidaPessoa()
	if errVal != nil {
		return nil, errVal
	}
	newPessoa.Imc = utils.CalculaImc(newPessoa.Peso, newPessoa.Altura)
	filtro := getFilter(nome)
	dados := bson.D{primitive.E{Key: "$set", Value: &newPessoa}}
	_, err := p.collection.UpdateOne(context.TODO(), filtro, dados)
	return newPessoa, err
}

// Exclui uma pessoa
func (p *Pessoa) Delete(nome string) (int64, error) {
	if nome == "" {
		return 0, errors.New("nome da pessoa não informado para exclusão")
	}
	filtro := getFilter(nome)
	result, err := p.collection.DeleteOne(context.TODO(), filtro)
	return result.DeletedCount, err
}

// Monta o filtro com o nome da pessoa para utilizar no CRUD
func getFilter(nome string) primitive.D {
	return bson.D{primitive.E{Key: "nome", Value: nome}}
}