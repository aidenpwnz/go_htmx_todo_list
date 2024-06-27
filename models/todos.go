package models

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

type TodoItem struct {
	Id          string
	Title       string
	Description string
}

func GenerateID() string {
	id := make([]byte, 16)
	_, err := rand.Read(id)
	if err != nil {
		log.Fatalf("Failed to generate ID: %v", err)
	}
	return hex.EncodeToString(id)
}
