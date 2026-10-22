package main

import (
	"crypto/hmac"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash"
	"strings"
)

type SessionData struct {
	Authenticated bool `json:"authenticated"`
}

type ErrSessionData struct {
	err error
}

func (err ErrSessionData) Error() string {
	return fmt.Sprintf("Session Data is invalid with error:\n %v", err.err)
}

func (sd *SessionData) ToBytes() ([]byte, error) {
	v, err := json.Marshal(sd)
	if err != nil {
		return nil, ErrSessionData{err}
	}
	return v, nil
}

type Session struct {
	key          []byte
	hashFunction func() hash.Hash
}

type (
	ErrKeySize          []int
	ErrSessionCookie    bool
	ErrInvalidSignature bool
)

func (e ErrKeySize) Error() string {
	return fmt.Sprintf("Invalid byte size for Key, possible sizes are: %v", []int(e))
}

func (e ErrSessionCookie) Error() string {
	return "Invalid cookie format."
}

func (e ErrInvalidSignature) Error() string {
	return "Invalid signature."
}

// NewSession creates a new Session with Key and Hash Function.
func NewSession(key string, hashFunction func() hash.Hash) (*Session, error) {
	keyByte := []byte(key)
	sizes := [3]int{16, 24, 32}
	for _, s := range sizes {
		if len(keyByte) == s {
			return &Session{keyByte, hashFunction}, nil
		}
	}
	return nil, ErrKeySize(sizes[:])
}

func (s *Session) Signature(data []byte) []byte {
	mac := hmac.New(s.hashFunction, s.key)
	mac.Write(data)
	return mac.Sum(nil)
}

func (s *Session) Create(data *SessionData) (string, error) {
	v, err := data.ToBytes()
	if err != nil {
		return "", err
	}
	signature := s.Signature(v)
	payload := base64.RawURLEncoding.EncodeToString(v)
	signatureBase64 := base64.RawURLEncoding.EncodeToString(signature)
	return fmt.Sprintf("%s.%s", payload, signatureBase64), nil
}

func (s *Session) Check(cookie string) (*SessionData, error) {
	cookieValues := strings.Split(cookie, ".")
	if len(cookieValues) != 2 {
		return nil, ErrSessionCookie(true)
	}
	payload, err := base64.RawURLEncoding.DecodeString(cookieValues[0])
	if err != nil {
		return nil, ErrSessionCookie(true)
	}
	signature, err := base64.RawURLEncoding.DecodeString(cookieValues[1])
	if err != nil {
		return nil, ErrSessionCookie(true)
	}
	if !hmac.Equal(signature, s.Signature(payload)) {
		return nil, ErrInvalidSignature(true)
	}
	var sessionData SessionData
	err = json.Unmarshal(payload, &sessionData)
	if err != nil {
		return nil, ErrSessionData{err}
	}
	return &sessionData, nil
}
