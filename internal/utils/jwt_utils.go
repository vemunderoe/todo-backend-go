package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
	"todo-backend/internal/models"
)

var secretKey = []byte("secret_key")

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Claims struct {
	User      models.User `json:"user"`
	ExpiresAt int64       `json:"exp"`
}

func GenerateJWT(user models.User) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	claims := Claims{
		User:      user,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	headerEncoded := base64.RawURLEncoding.EncodeToString(headerJSON)
	claimsEncoded := base64.RawURLEncoding.EncodeToString(claimsJSON)

	signature := createSignature(headerEncoded, claimsEncoded)

	token := strings.Join([]string{headerEncoded, claimsEncoded, signature}, ".")

	return token, nil
}

func ValidateJWT(token string) (*Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, ErrInvalidToken
	}

	headerEncoded, claimsEncoded, signature := parts[0], parts[1], parts[2]

	expectedSignature := createSignature(headerEncoded, claimsEncoded)
	if !hmac.Equal([]byte(signature), []byte(expectedSignature)) {
		return nil, ErrInvalidToken
	}

	// Decode and parse claims
	claimsJSON, err := base64.RawURLEncoding.DecodeString(claimsEncoded)
	if err != nil {
		return nil, err
	}

	var claims Claims
	if err := json.Unmarshal(claimsJSON, &claims); err != nil {
		return nil, err
	}

	// Check if the token has expired
	if time.Now().Unix() > claims.ExpiresAt {
		return nil, ErrExpiredToken
	}

	return &claims, nil
}

func createSignature(headerEncoded, claimsEncoded string) string {
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(headerEncoded + "." + claimsEncoded))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)
