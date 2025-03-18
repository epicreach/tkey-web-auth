package auth

import (
	"bytes"
	. "chalmers/tkey-group22/client/internal/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

// RemovePublicKey removes a public key from the authenticated user's account
// This requires that the app has the /api/remove-public-key endpoint
// It returns an error if the process fails
//
// Parameters:
// - appurl: The URL of the application server
// - username: The username of the authenticated user
// - label: The label of the public key to be removed
// - sessionCookie: The session cookie to be included in the request
//
// Returns:
// - An error if the process fails
func RemovePublicKey(appurl string, username string, label string, sessionCookie string) error {
	requestBody := RemovePublicKeyRequest{
		Label: label,
	}

	return sendRemovePublicKeyRequest(appurl, requestBody, sessionCookie)
}

// sendRemovePublicKeyRequest sends a request to remove a public key for a user
// It returns an error if the process fails
//
// Parameters:
// - appurl: The URL of the application server
// - requestBody: The request body containing the username and the label of the public key to be removed
// - sessionCookie: The session cookie to be included in the request
//
// Returns:
// - An error if the process fails
func sendRemovePublicKeyRequest(appurl string, requestBody RemovePublicKeyRequest, sessionCookie string) error {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	endpoint := "/api/remove-public-key"
	req, err := http.NewRequest("POST", appurl+endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", sessionCookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var responseBody map[string]string
		if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
			return fmt.Errorf("unexpected error: %s", resp.Status)
		}
		return fmt.Errorf("%s", responseBody["message"])
	}

	fmt.Printf("Public key removed successfully")
	return nil
}
