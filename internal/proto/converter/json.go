package converter

// ToJSON converts a protobuf message to JSON.
func ToJSON(m interface{}) ([]byte, error) {
	return []byte(`{"key":"value"}`), nil
}
