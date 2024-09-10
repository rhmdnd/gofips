// This example was lifted from https://pkg.go.dev/golang.org/x/crypto@v0.27.0/chacha20poly1305#example-NewX
package main

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
)

func main() {
	// key should be randomly generated or derived from a function like Argon2.
	key := make([]byte, chacha20poly1305.KeySize)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		panic(err)
	}

	// Encryption.
	var encryptedMsg []byte
	{
		msg := []byte("Gophers, gophers, gophers everywhere!")

		// Select a random nonce, and leave capacity for the ciphertext.
		nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(msg)+aead.Overhead())
		if _, err := rand.Read(nonce); err != nil {
			panic(err)
		}

		// Encrypt the message and append the ciphertext to the nonce.
		encryptedMsg = aead.Seal(nonce, nonce, msg, nil) // want "Seal is not a FIPS-validated implementation"
	}

	// Decryption.
	{
		if len(encryptedMsg) < aead.NonceSize() {
			panic("ciphertext too short")
		}

		// Split nonce and ciphertext.
		nonce, ciphertext := encryptedMsg[:aead.NonceSize()], encryptedMsg[aead.NonceSize():]

		// Decrypt the message and check it wasn't tampered with.
		plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", plaintext)
	}

}
