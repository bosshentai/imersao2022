package kafka

import (
	"encoding/json"
	route2 "github.com/bosshentai/imersao2022/simulador-aluno/application/route"
	"github.com/bosshentai/imersao2022/simulador-aluno/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"time"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}

	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaReadTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
