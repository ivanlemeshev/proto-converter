package converter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ivanlemeshev/proto-converter/internal/proto/converter"
)

type testMessage struct {
	Key string
}

func TestToJSON(t *testing.T) {
	expected := []byte(`{"key":"value"}`)
	actual, err := converter.ToJSON(&testMessage{
		Key: "value",
	})

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
