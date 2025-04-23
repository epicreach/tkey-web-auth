package structs

// RegisterRequest represents the data required to register a new user
// It includes the username, the user's public key, and the label for the public key
type RegisterRequest struct {
	Username string `json:"username"`
	Pubkey   []byte `json:"pubkey"`
	Label    string `json:"label"`
}

// UnregisterRequest represents the payload for a unregister request
// It contains the username of the user attempting to unregister
type UnregisterRequest struct {
	Username string `json:"username"`
}

// LoginRequest represents the payload for a login request
// It contains the username of the user attempting to log in
type LoginRequest struct {
	Username string `json:"username"`
}

// LoginResponse represents the response received after a login attempt
// It contains a challenge string and a signature string, both of which are used to verify the authenticity of the login request
type LoginResponse struct {
	Challenge string `json:"challenge"`
	Signature string `json:"signature"`
}

// VerifyRequest represents a request to verify a user's identity
// It contains the username of the user and a cryptographic signature to authenticate the request
type VerifyRequest struct {
	Username  string `json:"username"`
	Signature []byte `json:"signature"`
}

type VerifyResponse struct {
}

type GetAndSignResponse struct {
	User            string `json:"user"`
	SignedChallenge []byte `json:"signed_challenge"`
}

// AddPublicKeyRequest represents a response to add a new public key for a user
// It contains the new public key to be added
type AddPublicKeyResponse struct {
	Pubkey []byte `json:"pubkey"`
}
