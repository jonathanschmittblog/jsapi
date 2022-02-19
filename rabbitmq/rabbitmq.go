package rabbitmq

import (
	"jsapi/utils"
	"os"

	"github.com/streadway/amqp"
)

type RabbitMq struct {
	Connection *amqp.Connection
	Channel *amqp.Channel
	Queue *amqp.Queue
}

// Método público para enviar a mensagem à fila RabbitMQ
func SendMessage(nomeQueue string, msg string) {
	queue := new(nomeQueue)
	queue.send(msg)
	defer queue.Connection.Close()
	defer queue.Channel.Close()
}

// Cria um novo objeto de RabbitMq
func new(nome string) *RabbitMq {
	rabbitMq := &RabbitMq{}
	// Conecta no server RabbitMQ
	rabbitMq.connect()
	rabbitMq.createChannel()
	rabbitMq.createQueue(nome)
	return rabbitMq
}

func (r *RabbitMq) connect() {
	conn, err := amqp.Dial(os.Getenv("JSAPI_RABBITMQ_DIAL"))
	utils.FailOnError(err, "Falha ao conectar no servidor RabbitMQ.")
	r.Connection = conn
}

// Cria o canal de comunicação com a API
func (r *RabbitMq) createChannel() {
	ch, err := r.Connection.Channel()
	utils.FailOnError(err, "Falha ao criar canal no servidor RabbitMQ.")
	r.Channel = ch
}

// Cria o canal de comunicação com a API
func (r *RabbitMq) createQueue(nome string) {
	// Envia mensagem para a fila
	queue, err := r.Channel.QueueDeclare(
		nome, 
		false,   
		false,   
		false,   
		false,   
		nil,     
	)
	utils.FailOnError(err, "Falha ao criar a fila.")
	r.Queue = &queue
}

// Envia uma mensagem para a fila RabbitMQ
func (r *RabbitMq) send(message string) {
	body := message
	err := r.Channel.Publish(
		"",     
		r.Queue.Name, 
		false,  
		false,  
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	utils.FailOnError(err, "Falha ao publicar a mensagem.")
}
