package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func Decrypt(key []byte, text []byte) ([]byte, error) {

	// Init decipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if (len(text) % aes.BlockSize) != 0 {
		return nil, errors.New("wrong blocksize")
	}

	// Getting the IV
	iv := text[:aes.BlockSize]

	// Actual decryption
	decodedCipherMsg := text[aes.BlockSize:]
	cfbDecrypter := cipher.NewCFBDecrypter(block, iv)
	cfbDecrypter.XORKeyStream(decodedCipherMsg, decodedCipherMsg)

	// Removing Padding
	length := len(decodedCipherMsg)
	paddingLen := int(decodedCipherMsg[length-1])
	result := decodedCipherMsg[:(length - paddingLen)]
	return result, nil
}
