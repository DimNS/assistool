package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

// CryptTab struct.
type CryptTab struct{}

// NewCryptTab creates a new CryptTab.
func NewCryptTab() *CryptTab {
	return &CryptTab{}
}

func (*CryptTab) Encrypt(secretKey string, decryptedString string, withGCM bool) (string, error) { //nolint:revive // it's ok
	plainText := []byte(decryptedString)

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("new cipher: %v", err)
	}

	if withGCM {
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			return "", fmt.Errorf("new GCM: %v", err)
		}

		nonce := make([]byte, gcm.NonceSize())
		if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
			return "", fmt.Errorf("read nonce: %v", err)
		}

		cipherText := gcm.Seal(nonce, nonce, plainText, nil)

		return base64.RawStdEncoding.EncodeToString(cipherText), nil
	}

	// Make the cipher text a byte array of size BlockSize + the length of the message
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	// iv is the ciphertext up to the blocksize (16)
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("read iv: %v", err)
	}

	// Encrypt the data:
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.RawStdEncoding.EncodeToString(cipherText), err
}

func (*CryptTab) Decrypt(secretKey string, encryptedString string, withGCM bool) (string, error) { //nolint:revive // it's ok
	cipherText, err := base64.RawStdEncoding.DecodeString(encryptedString)
	if err != nil {
		return "", fmt.Errorf("base64 decode: %v", err)
	}

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("new cipher: %v", err)
	}

	if withGCM {
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			return "", fmt.Errorf("new GCM: %v", err)
		}

		// Extract nonce from the beginning of the ciphertext
		nonceSize := gcm.NonceSize()
		if len(cipherText) < nonceSize {
			return "", errors.New("ciphertext too short")
		}

		nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:] //nolint:govet // it's ok

		// Decrypt and verify MAC
		plainText, err := gcm.Open(nil, nonce, cipherText, nil)
		if err != nil {
			return "", fmt.Errorf("decryption failed: %v", err)
		}

		return string(plainText), nil
	}

	// IF the length of the cipherText is less than 16 Bytes:
	if len(cipherText) < aes.BlockSize {
		return "", errors.New("ciphertext block size is too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// Decrypt the message
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), err
}
