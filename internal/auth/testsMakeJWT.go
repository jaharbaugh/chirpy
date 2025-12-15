package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
)

func TestMakeJWT_Success(t *testing.T) {
	userID := uuid.New()
	secret := "super-secret-key"
	expiresIn := time.Minute

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if tokenString == "" {
		t.Fatal("expected token string, got empty string")
	}
}

func TestMakeJWT_ClaimsAreCorrect(t *testing.T) {
	userID := uuid.New()
	secret := "another-secret"
	expiresIn := time.Minute

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	parsedToken, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	if err != nil {
		t.Fatalf("failed to parse token: %v", err)
	}

	claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims)
	if !ok {
		t.Fatal("claims were not of type RegisteredClaims")
	}

	if claims.Issuer != "chirpy" {
		t.Errorf("expected issuer 'chirpy', got %q", claims.Issuer)
	}

	if claims.Subject != userID.String() {
		t.Errorf("expected subject %q, got %q", userID.String(), claims.Subject)
	}

	if claims.IssuedAt == nil {
		t.Error("expected IssuedAt to be set")
	}

	if claims.ExpiresAt == nil {
		t.Error("expected ExpiresAt to be set")
	}
}

func TestMakeJWT_ExpirationIsInFuture(t *testing.T) {
	userID := uuid.New()
	secret := "expiration-test"
	expiresIn := 2 * time.Minute

	before := time.Now().UTC()

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	parsedToken, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	if err != nil {
		t.Fatalf("failed to parse token: %v", err)
	}

	claims := parsedToken.Claims.(*jwt.RegisteredClaims)

	if claims.ExpiresAt.Time.Before(before) {
		t.Error("expected ExpiresAt to be in the future")
	}
}

func TestMakeJWT_InvalidSecretFailsValidation(t *testing.T) {
	userID := uuid.New()
	secret := "correct-secret"
	wrongSecret := "wrong-secret"

	tokenString, err := MakeJWT(userID, secret, time.Minute)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(wrongSecret), nil
		},
	)

	if err == nil {
		t.Fatal("expected error when parsing with wrong secret, got nil")
	}
}
