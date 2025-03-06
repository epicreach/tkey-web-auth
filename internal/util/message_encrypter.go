package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

// encryptMessage encrypts a given message
// It uses RSA-OAEP with SHA-256 hashing for encryption
//
// Parameters:
// message: plaintext of message to be encrypted with type []byte
// public_key: the RSA public key
//
// Returns:
// Returns encrypted message with type []byte
func encryptMessage(message []byte, publicKey *rsa.PublicKey) []byte {
	hash := sha256.New()
	random := rand.Reader
	encrypted_message, err := rsa.EncryptOAEP(hash, random, publicKey, message, nil)
	if err != nil {
		panic(err)
	}

	return encrypted_message
}

// decryptMessage decrypts a given ciphertext
// It uses RSA-OAEP with SHA-256 hashing for decryption
//
// Parameters:
// ciphertext: encrypted message to be decrypted with type []byte
// private_key: the RSA private key
//
// Returns:
// Returns decrypted message with type []byte
func decryptMessage(ciphertext []byte, privateKey *rsa.PrivateKey) []byte {
	hash := sha256.New()
	random := rand.Reader
	decrypted_message, err := rsa.DecryptOAEP(hash, random, privateKey, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	return decrypted_message
}
