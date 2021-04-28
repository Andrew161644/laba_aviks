package connection

import (
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

func HandleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Connect(connString string, queueName string) *Connection {
	conn, err := amqp.Dial(connString)
	HandleError(err, "Can't connect to AMQP")
	amqpChannel, err := conn.Channel()
	HandleError(err, "Can't create a amqpChannel")
	queue, err := amqpChannel.QueueDeclare(queueName, true, false, false, false, nil)
	HandleError(err, "Could not declare `add` queue")
	err = amqpChannel.Qos(1, 0, false)
	HandleError(err, "Could not configure QoS")

	return &Connection{
		Connection: conn,
		Channel:    amqpChannel,
		Queue:      queue,
	}
}

type Connection struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

func (connection *Connection) DeleteQueue() {
	defer connection.Connection.Close()
	defer connection.Channel.Close()
	connection.Channel.QueueDelete(connection.Queue.Name, true, true, true)
}

func CreateInOutConnectionForUser(connectionString string, queueName string, user models.UserModel) (Connection, Connection) {
	var connectionIn = *Connect(connectionString, queueName)
	var connectionOut = *Connect(connectionString, queueName+"_"+strconv.Itoa(user.ID))
	return connectionIn, connectionOut
}
