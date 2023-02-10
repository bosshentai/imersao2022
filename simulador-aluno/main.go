package main

import (
	"fmt"
	kafka2 "github.com/bosshentai/imersao2022/simulador-aluno/application/kafka"
	"github.com/bosshentai/imersao2022/simulador-aluno/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
func main() {

	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		go kafka2.Produce(msg)
		fmt.Println(string(msg.Value))
	}

	//producer := kafka.NewKafkaProducer()
	//kafka.Publish("ola", "readtest", producer)

	//for {
	//	_ = 1
	//}

	//route := route2.Route{
	//	ID:       "1",
	//	ClientID: "1",
	//	//Positions: nil,
	//}
	//
	//route.LoadPositions()
	//stringjson, _ := route.ExportJsonPositions()
	//fmt.Println(stringjson[0])
}
