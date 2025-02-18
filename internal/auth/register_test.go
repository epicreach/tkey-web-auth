package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"
)

func TestConvertToJSON(t *testing.T) {
	data := map[string]string{"username": "test_username", "pubkey": "my_pubkey"}

	result := convertToJSON(data)

	expectedResult, err := json.Marshal(data)
	if err != nil {
		t.Fatal("Error marshaling expected result:", err)
	}

	if string(result) != string(expectedResult) {
		t.Errorf("Expected %s, but got %s", expectedResult, result)
	}
}

func TestCreateRequest(t *testing.T) {

	jsonData := []byte(`{"username":"test_username", "pubkey":"my_pubkey"}`)
	url := "http://localhost:8080/api/endpoint"

	req := createRequest(jsonData, url)

	if req.Method != "POST" {
		t.Errorf("Expected method 'POST', but got '%s'", req.Method)
	}

	if req.URL.String() != url {
		t.Errorf("Expected URL '%s', but got '%s'", url, req.URL.String())
	}

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		t.Errorf("Failed to read request body: %s", err)
	}

	if !bytes.Equal(bodyBytes, jsonData) {
		t.Errorf("Expected body %s, but got %s", jsonData, bodyBytes)
	}
}
