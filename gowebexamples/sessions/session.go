package main

import (
	"crypto/hmac"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
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
	ErrSessionCookie struct {
		err error
	}
	ErrInvalidSignature bool
)

func (e ErrSessionCookie) Error() string {
	return fmt.Sprintf("Invalid cookie: %v", e.err)
}

func (e ErrInvalidSignature) Error() string {
	return "Invalid signature."
}

type KeyEncoding int

const (
	EncodingBase64 KeyEncoding = iota // 0
	EncodingHex                       // 1
)

// NewSession creates a new Session with Key and Hash Function.
func NewSession(key string, encoding KeyEncoding, hashFunction func() hash.Hash) (*Session, error) {
	var rawKey []byte
	var err error

	switch encoding {
	case EncodingBase64:
		rawKey, err = base64.StdEncoding.DecodeString(key)
	case EncodingHex:
		rawKey, err = hex.DecodeString(key)
	default:
		return nil, fmt.Errorf("unsupported key encoding format")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to decode key: %w", err)
	}

	sizes := [3]int{16, 24, 32}
	for _, s := range sizes {
		if len(rawKey) == s {
			return &Session{rawKey, hashFunction}, nil
		}
	}
	return nil, fmt.Errorf("invalid byte size for Key, possible sizes are: %v", sizes[:])
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
		return nil, ErrSessionCookie{errors.New("cookie format is invalid")}
	}
	payload, err := base64.RawURLEncoding.DecodeString(cookieValues[0])
	if err != nil {
		return nil, ErrSessionCookie{err}
	}
	signature, err := base64.RawURLEncoding.DecodeString(cookieValues[1])
	if err != nil {
		return nil, ErrSessionCookie{err}
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
