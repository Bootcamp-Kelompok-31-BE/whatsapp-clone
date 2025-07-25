package middlewares

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/constant"
)

func Encrypt(data []byte) ([]byte, error) {
	key := constant.SecretKeyE2EE
	// create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// menampung ke chipertext dimana data yang telah diukur panjangannya dan bergabung dengan block size
	cipherdata := make([]byte, aes.BlockSize+len(data))
	iv := cipherdata[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// encrypt the data
	encrypting := cipher.NewCFBEncrypter(block, iv)
	encrypting.XORKeyStream(cipherdata[aes.BlockSize:], data)

	// Return base64-encoded ciphertext
	encoded := []byte(base64.StdEncoding.EncodeToString(cipherdata))
	return encoded, nil

}

