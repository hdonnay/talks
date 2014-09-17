// +build nocompile

package main

type (
	// Header contains information about the encryption scheme
	Header struct {
		Scheme Scheme
		IV     []byte
		MAC    []byte
	}
	// Key is used for symmetric encryption in/out of the blob store
	Key struct {
		Scheme Scheme
		// Symetric cipher key
		Key []byte
		// Key for HMAC
		HMAC []byte
	}
	// Encryption scheme
	Scheme int
)
