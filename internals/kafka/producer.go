package producer

import (
	"log"
	"os"

	"github.com/ozonva/ova-book-api/internals/entities/book"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type kafkaProducer struct {
	producer *kafka.Producer
}

func New(kafkaServers string) kafkaProducer {
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatalf("Ошибка при получении имени хоста: %v", err)
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServers,
		"client.id":         hostName,
		"acks":              "all",
	})
	if err != nil {
		log.Fatalf("Ошибка при подключении к Kafka: %v", err)
	}

	return kafkaProducer{producer: p}
}

func (p *kafkaProducer) produceMessage(topic, bookISBN10 string, deliveryChan chan kafka.Event) error {
	err := p.producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(bookISBN10),
		},
		deliveryChan,
	)
	return err
}

func (p *kafkaProducer) CreateEvent(createdBook book.Book) error {
	deliveryChan := make(chan kafka.Event, 1000)
	topic := "created"
	err := p.produceMessage(topic, createdBook.ISBN10, deliveryChan)
	return err
}

func (p *kafkaProducer) UpdateEvent(updatedBook book.Book) error {
	deliveryChan := make(chan kafka.Event, 1000)
	topic := "updated"
	err := p.produceMessage(topic, updatedBook.ISBN10, deliveryChan)
	return err
}

func (p *kafkaProducer) DeleteEvent(deletedBook book.Book) error {
	deliveryChan := make(chan kafka.Event, 1000)
	topic := "deleted"
	err := p.produceMessage(topic, deletedBook.ISBN10, deliveryChan)
	return err
}
