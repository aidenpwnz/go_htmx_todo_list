package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"todo_list/models"
)

func TestGenerateID(t *testing.T) {
	id := models.GenerateID()
	assert.Len(t, id, 32)
	assert.NotEmpty(t, id)
}

func TestGenerateIDUniqueness(t *testing.T) {
	ids := make(map[string]struct{})
	for i := 0; i < 1000; i++ {
		id := models.GenerateID()
		_, exists := ids[id]
		assert.False(t, exists, "Generated ID should be unique")
		ids[id] = struct{}{}
	}
}
