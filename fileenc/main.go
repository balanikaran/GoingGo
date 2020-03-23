package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	EncryptFile("normal.txt", []byte("mypasswordfile"))
	err := DecryptFile("normal.txt.enc", []byte("mypasswordfile"))
	fmt.Println(err)
}

// EncryptFile -
func EncryptFile(source string, enckey []byte) error {

	if _, err := os.Stat(source); os.IsNotExist(err) {
		return errors.New("file does not exist")
	}

	plainText, err := ioutil.ReadFile(source)
	if err != nil {
		return errors.New("error reading the file")
	}

	// make nonce using the
	key := enckey
	nonce := make([]byte, 12)

	// randomize the nonce
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return errors.New("error randomizing nonce")
	}

	// creating derived key
	iterationCount := 4096
	derivedKeyLength := 32
	derivedKey := pbkdf2.Key(key, nonce, iterationCount, derivedKeyLength, sha1.New)

	// creating a block cipher using derived key
	aesBlock, err := aes.NewCipher(derivedKey)
	if err != nil {
		return errors.New("error creating aes block cipher")
	}

	// creating gcm block cipher
	gcmBlock, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return errors.New("error creating gcm block cipher")
	}

	// creating cipherText using gcm block
	cipherText := gcmBlock.Seal(nil, nonce, plainText, nil)

	// appending nonce at the end of cipherText
	cipherText = append(cipherText, nonce...)

	// creating new encrypted file
	eFile, err := os.Create(source + ".enc")
	if err != nil {
		return errors.New("error creating the new enc file")
	}

	// copy encypted text to new encrypted file
	_, err = io.Copy(eFile, bytes.NewReader(cipherText))
	if err != nil {
		return errors.New("error writing to the enc file")
	}

	return nil
}

// DecryptFile -
func DecryptFile(source string, decKey []byte) error {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return errors.New("source file does not exist")
	}

	cipherText, err := ioutil.ReadFile(source)
	if err != nil {
		return errors.New("error reading encrypted file")
	}

	// create salt using key
	key := decKey
	saltHex := cipherText[len(cipherText)-12:]
	saltString := hex.EncodeToString(saltHex)

	// create nonce using salt
	nonce, err := hex.DecodeString(saltString)
	if err != nil {
		return errors.New("error reading encrypted file")
	}

	// create derived key
	iterationCount := 4096
	derivedKeyLength := 32
	derivedKey := pbkdf2.Key(key, nonce, iterationCount, derivedKeyLength, sha1.New)

	// creating block cipher using derived key
	aesBlock, err := aes.NewCipher(derivedKey)
	if err != nil {
		return errors.New("error creating aes block")
	}

	// create gcm block
	gcmBlock, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return errors.New("error creating gcm block")
	}

	// creating plaintext back
	plainText, err := gcmBlock.Open(nil, nonce, cipherText[:len(cipherText)-12], nil)
	if err != nil {
		return errors.New("error decrypting file")
	}

	// creating decrypted file
	dFile, err := os.Create(source + ".txt1")
	if err != nil {
		return errors.New("error creating decrypted file")
	}

	// copy plain text to decrypted file
	_, err = io.Copy(dFile, bytes.NewReader(plainText))
	if err != nil {
		return errors.New("error writing decrypted file")
	}

	return nil
}
