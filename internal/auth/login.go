package auth

import (
	"bytes"
	. "chalmers/tkey-group22/internal/structs"
	. "chalmers/tkey-group22/internal/tkey"
	"encoding/json"
	"fmt"
	"net/http"
)

func Login() error {
	username := getUsername()

	c := &http.Client{}

	challengeResponse, err := getChallenge(username)
	if err != nil {
		return err
	}

	// TODO: Implement signature verification
	// if !verifySignature(challengeResponse) {
	// 	return fmt.Errorf("signature verification failed")
	// }

	signedChallenge, err := signChallenge(username, challengeResponse)
	if err != nil {
		return err
	}

	baseUrl := "http://localhost:8080"
	endpoint := "/api/verify"

	body, err := json.Marshal(signedChallenge)
	if err != nil {
		return err
	}

	resp, err := c.Post(baseUrl+endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("error in response when sending signature")
	}

	return nil
}

func signChallenge(username string, challenge *LoginResponse) (*VerifyRequest, error) {
	// Sign the challenge
	sig, err := Sign([]byte(challenge.Challenge))
	if err != nil {
		return nil, err
	}

	return &VerifyRequest{
		Username:  username,
		Signature: string(sig),
	}, nil
}

func getChallenge(user string) (*LoginResponse, error) {
	baseUrl := "http://localhost:8080"
	endpoint := "/api/login"

	// Get the signature and message from the endpoint

	c := &http.Client{}

	body, err := json.Marshal(LoginRequest{Username: user})
	if err != nil {
		return nil, err
	}

	resp, err := c.Post(baseUrl+endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error in response when requesting challenge")
	}

	var res *LoginResponse
	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, fmt.Errorf("error decoding challenge response")
	}

	return res, nil
}

// TODO: Uncomment once server verification is implemented
// func verifySignature(res *LoginResponse) bool {
// 	pubkey, err := getPublicKey()
// 	if err != nil {
// 		return false
// 	}
// 	return ed25519.Verify(pubkey, []byte(res.Challenge), []byte(res.Signature))
// }

// func getPublicKey() ([]byte, error) {
// 	baseUrl := "http://localhost:8080"
// 	endpoint := "/api/public"

// 	c := &http.Client{}

// 	resp, err := c.Get(baseUrl + endpoint)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if resp.StatusCode != 200 {
// 		return nil, fmt.Errorf("error in response when getting public key")
// 	}

// 	var pubkey []byte

// 	err = json.NewDecoder(resp.Body).Decode(&pubkey)

// 	return pubkey, err
// }
