package auth

import (
	"bytes"
	. "chalmers/tkey-group22/client/internal/structs"
	"chalmers/tkey-group22/client/internal/tkey"
	"encoding/json"
	"fmt"
	"net/http"
)

// AddPublicKey adds a new public key to the authenticated user's account
// This requires that the app has the /api/add-public-key endpoint
// It returns an error if the process fails
//
// Parameters:
// - appurl: The URL of the application server
// - username: The username of the authenticated user
// - label: The label for the new public key
// - sessionCookie: The session cookie to be included in the request
//
// Returns:
// - An error if the process fails
func AddPublicKey(appurl string, username string, label string, sessionCookie string) error {
	pubkey, err := tkey.GetTkeyPubKey()
	if err != nil {
		return err
	}

	requestBody := AddPublicKeyRequest{
		Pubkey: []byte(pubkey),
		Label:  label,
	}

	return sendAddPublicKeyRequest(appurl, requestBody, sessionCookie)
}

// sendAddPublicKeyRequest sends a request to add a new public key for a user
// It returns an error if the process fails
//
// Parameters:
// - appurl: The URL of the application server
// - requestBody: The request body containing the username and the new public key
//
// Returns:
// - An error if the process fails
func sendAddPublicKeyRequest(appurl string, requestBody AddPublicKeyRequest, sessionCookie string) error {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	endpoint := "/api/add-public-key"
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

	fmt.Printf("Public key added successfully")
	return nil
}
