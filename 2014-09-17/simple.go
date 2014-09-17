// +build nocompile

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"log"
)

// bad example OMIT
// Example using AES
func One(plaintext io.Reader, key []byte) (io.Reader, error) {
	blk, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, blk.BlockSize())                                      // HLcallout
	return cipher.StreamReader{S: cipher.NewOFB(blk, iv), R: plaintext}, nil // HLcallout
}

// good example OMIT
// Example using AES
func Two(plaintext io.Reader, key []byte) (io.Reader, error) {
	blk, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, blk.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil { // HLcallout
		return nil, err
	}
	s := cipher.StreamReader{S: cipher.NewOFB(blk, iv), R: plaintext}
	r, w := io.Pipe()
	go func() {
		defer w.Close()
		h := hmac.New(sha256.New, key) // HLcallout
		w.Write(iv)                    // HLcallout
		t := io.TeeReader(s, h)
		if _, err := io.Copy(w, t); err != nil {
			log.Println("on no :(")
		}
		w.Write(h.Sum(nil)) // HLcallout
	}()
	return r, nil
}
