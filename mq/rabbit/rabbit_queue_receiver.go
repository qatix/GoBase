package rabbit

import (
	"log"
	"fmt"
	"github.com/streadway/amqp"
)

func failOnError(err error,msg string)  {
	if err != nil{
		log.Fatalf("%s:%s",msg,err)
		panic(fmt.Sprintf("%s:%s",msg,err))
	}
}

func main(){
	conn,err := amqp.Dial("amqp://guest:guest@192.168.1.100:5672")
	failOnError(err,"Failed to connect to RabbitMQ")

	defer conn.Close()

	ch,err := conn.Channel()
	failOnError(err,"Failed to open a channel")
	defer ch.Close()

	q,err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err,"Failed to declare a queue")

	//func (ch *Channel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args Table) (<-chan Delivery, error)
	msgs,err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err,"Failed to register a consumer")

	forever := make(chan bool) //构造一个永久等待
	go func() {
		for d:= range msgs{
			log.Printf("Received a message:%s",d.Body)
		}
	}()

	log.Printf(" [*] Waiting for message,to exit press ctrl+c")
	<-forever //没有值写入，永久等待
}
