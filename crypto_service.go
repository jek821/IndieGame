package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privKey, &privKey.PublicKey, nil
}

// GenerateSessionKey returns a 32-byte random key for AES-256.
func GenerateSessionKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	return key, err
}

// EncryptSessionKey is called by the client to encrypt the session key with the server's public key.
func EncryptSessionKey(pubKey *rsa.PublicKey, sessionKey []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, sessionKey, nil)
}

// DecryptSessionKey is called by the server to recover the session key using its private key.
func DecryptSessionKey(privKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, privKey, ciphertext, nil)
}
