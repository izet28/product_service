package events

import (
	"encoding/json"
	"log"
	"product_service/models"

	"github.com/IBM/sarama"
)

func PublishProductCreatedEvent(product *models.Product) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Tentukan broker Kafka (sesuaikan dengan konfigurasi Kafka)
	brokers := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Gagal terhubung ke Kafka: %v", err)
		return err
	}
	defer producer.Close()

	// Encode product ke JSON
	productData, err := json.Marshal(product)
	if err != nil {
		return err
	}

	// Buat pesan Kafka
	msg := &sarama.ProducerMessage{
		Topic: "product-events", // Kafka topic yang akan digunakan
		Key:   sarama.StringEncoder(product.ID),
		Value: sarama.ByteEncoder(productData),
	}

	// Kirim pesan
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Produk berhasil dipublikasikan ke Kafka pada partition %d, offset %d", partition, offset)
	return nil
}
