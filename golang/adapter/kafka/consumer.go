package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Consumer struct {
	c *kafka.Consumer
}

type Config struct {
	Broker  string
	Topic   string
	GroupID string
	// SRURL   string
}

// CDCEvent sesuai struktur Debezium
type CDCEvent struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	Before *Product `json:"before"`
	After  *Product `json:"after"`
	Op     string   `json:"op"`
}

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     string  `json:"price"` // nanti bisa decode Base64 ke decimal
	Stock     int64   `json:"stock"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}

func NewConsumer(conf Config) (*Consumer, error) {
	// Buat consumer
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": conf.Broker,      // broker Kafka
		"group.id":          conf.GroupID,     // ID consumer group
		"auto.offset.reset": "earliest",       // baca dari awal kalau group baru
		"isolation.level":   "read_committed", // baca hanya pesan yang sudah commit → penting untuk EOS
	})
	if err != nil {
		log.Fatalf("NewConsumer:%v", err)
	}

	// Subscribe topic
	if err := c.SubscribeTopics([]string{conf.Topic}, nil); err != nil {
		log.Fatalf("SubscribeTopics:%v", err)
	}

	return &Consumer{c: c}, nil
}

func (cs *Consumer) Consume(ctx context.Context) {
	log.Println("Consume started...")
	defer cs.c.Close()

	index := 0
	for {
		select {
		case <-ctx.Done():
			log.Println("Shutting down consumer...")
			return
		default:
			index++
			msg, err := cs.c.ReadMessage(-1)
			if err != nil {
				log.Printf("Consume error [%d]: %v", index, err)
				continue
			}

			// log.Printf("Received[%d]: %s", index, string(msg.Value))

			var event CDCEvent
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				log.Printf("JSON error: %v", err)
				continue
			}

			// kirim ke handler
			cs.handleEvent(event)
		}
	}
}

func (cs *Consumer) handleEvent(e CDCEvent) {
	js, _ := json.MarshalIndent(e, "", "  ")
	log.Printf("Got event: %s", string(js))
	// TODO: integrasi ke DB, API, dsb.
}
