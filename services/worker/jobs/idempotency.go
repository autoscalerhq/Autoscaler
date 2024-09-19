package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"log"
)

// CheckUnique checks if the UUID already exists in JetStream KeyValue store
func CheckUnique(idempotencyKV jetstream.KeyValue, uuid string) (bool, error) {
	ctx := context.Background()
	_, err := idempotencyKV.Get(ctx, uuid)

	if err.Error() == nats.ErrKeyNotFound.Error() {
		return true, nil // UUID does not exist, hence it's unique
	} else if err != nil {
		return false, err // Something went wrong
	}

	return false, nil // UUID exists
}

func GetIdempotencyKey(idemptencyKV jetstream.KeyValue) (string, error) {
	var unique bool
	var newUUID uuid.UUID

	// Loop until a unique UUID is generated
	for {
		tempUUID, err := uuid.NewV7()
		if err != nil {
			log.Fatalf("Failed to generate UUID: %v", err)
		}

		unique, err = CheckUnique(idemptencyKV, newUUID.String())
		if err != nil {
			log.Fatalf("Failed to check UUID uniqueness: %v", err)
			return "", err
		}

		if unique {
			newUUID = tempUUID
			ctx := context.Background()
			_, err := idemptencyKV.Create(ctx, newUUID.String(), []byte{})
			if err != nil {
				log.Fatalf("Failed to set UUID for uniqueness: %v", err)
				return "", err
			}
			break
		}
	}

	return newUUID.String(), nil
}

func addIdempotencyKey(data []byte, idempotencyKey string) ([]byte, string, error) {
	// Unmarshal the data into a map
	var messageMap map[string]interface{}
	if err := json.Unmarshal(data, &messageMap); err != nil {
		return nil, "", fmt.Errorf("failed to unmarshal data: %w", err)
	}

	subject := messageMap["newSubject"].(string)

	delete(messageMap, "newSubject")
	// Add the idempotency key
	messageMap["idempotency"] = idempotencyKey

	// Marshal the modified map back to a JSON byte array
	modifiedData, err := json.Marshal(messageMap)
	if err != nil {
		return nil, "", fmt.Errorf("failed to marshal modified data: %w", err)
	}

	return modifiedData, subject, nil
}
