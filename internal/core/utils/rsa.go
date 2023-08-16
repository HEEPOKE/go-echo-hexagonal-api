package utils

// import (
// 	"crypto/rand"
// 	"crypto/rsa"
// 	"crypto/sha1"
// 	"crypto/x509"
// 	"encoding/base64"
// 	"encoding/pem"
// 	"errors"
// 	"fmt"

// 	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/config"
// )

// var privateKeyPem = []byte(`-----BEGIN RSA PRIVATE KEY-----` + config.Cfg.PRIVATE_KEY + `-----END RSA PRIVATE KEY-----`)

// var publicKeyPem = []byte(`-----BEGIN PUBLIC KEY-----` + config.Cfg.PUBLIC_KEY + `-----END PUBLIC KEY-----`)

// func Decrypt(encryptedMessage string) (string, error) {
// 	encryptedBuffer, err := base64.StdEncoding.DecodeString(encryptedMessage)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to decode encrypted message: %w", err)
// 	}

// 	block, _ := pem.Decode(privateKeyPem)
// 	if block == nil {
// 		return "", errors.New("failed to parse private key")
// 	}

// 	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to parse private key: %w", err)
// 	}

// 	decryptedBuffer, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, privateKey, encryptedBuffer, nil)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to decrypt message: %w", err)
// 	}

// 	return string(decryptedBuffer), nil
// }

// func Encrypt(plainText string) (string, error) {
// 	plainTextBuffer := []byte(plainText)

// 	block, _ := pem.Decode(publicKeyPem)
// 	if block == nil {
// 		return "", errors.New("failed to parse public key")
// 	}

// 	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to parse public key: %w", err)
// 	}

// 	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
// 	if !ok {
// 		return "", errors.New("invalid public key type")
// 	}

// 	encryptedBuffer, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, rsaPublicKey, plainTextBuffer, nil)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to encrypt message: %w", err)
// 	}

// 	encryptedMessage := base64.StdEncoding.EncodeToString(encryptedBuffer)
// 	return encryptedMessage, nil
// }
