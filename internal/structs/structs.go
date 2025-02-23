package structs

type RegisterRequest struct {
	Username string `json:"username"`
	Pubkey   []byte `json:"pubkey"`
}

type LoginRequest struct {
	Username string `json:"username"`
}

type LoginResponse struct {
	Challenge string `json:"challenge"`
	Signature string `json:"signature"`
}

type VerifyRequest struct {
	Username  string `json:"username"`
	Signature []byte `json:"signature"`
}

type VerifyResponse struct {
}
