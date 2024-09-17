package jobs

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"log"
)

//messageStruct := Message{
//	Idempotency: "unique-id-1234", // Replace with actual idempotency value
//	NextSubject: "job.client",
//}
//
//messageBytes, err := json.Marshal(messageStruct)
//if err != nil {
//	println("Failed to serialize message: %v", err)
//}
//
//messageString := string(messageBytes)

type IdempotentMessage struct {
	Idempotency string `json:"idempotency"`
	Message     string `json:"message"`
}

// CheckUnique checks if the UUID already exists in JetStream KeyValue store
func CheckUnique(idempotencyKV jetstream.KeyValue, uuid string) (bool, error) {
	ctx := context.Background()
	_, err := idempotencyKV.Get(ctx, uuid)
	if errors.Is(err, nats.ErrKeyNotFound) {
		return true, nil // UUID does not exist, hence it's unique
	} else if err != nil {
		return false, err // Something went wrong
	}
	return false, nil // UUID exists
}

func GetIdempotencyKey(idemptencyKV *jetstream.KeyValue) (string, error) {
	var unique bool
	var newUUID uuid.UUID

	// Loop until a unique UUID is generated
	for {
		newUUID, err := uuid.NewV7()
		if err != nil {
			log.Fatalf("Failed to generate UUID: %v", err)
		}

		unique, err = CheckUnique(*idemptencyKV, newUUID.String())
		if err != nil {
			log.Fatalf("Failed to check UUID uniqueness: %v", err)
			return "", err
		}

		if unique {
			break
		}
	}

	return newUUID.String(), nil
}
