package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
)

func TestValidateJWT_Success(t *testing.T) {
	userID := uuid.New()
	secret := "valid-secret"

	token, err := MakeJWT(userID, secret, time.Minute)
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}

	parsedID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if parsedID != userID {
		t.Errorf("expected userID %v, got %v", userID, parsedID)
	}
}

func TestValidateJWT_InvalidSignature(t *testing.T) {
	userID := uuid.New()
	secret := "correct-secret"
	wrongSecret := "wrong-secret"

	token, err := MakeJWT(userID, secret, time.Minute)
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}

	_, err = ValidateJWT(token, wrongSecret)
	if err == nil {
		t.Fatal("expected error for invalid signature, got nil")
	}
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
	userID := uuid.New()
	secret := "expired-secret"

	// Token already expired
	token, err := MakeJWT(userID, secret, -time.Minute)
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}

	_, err = ValidateJWT(token, secret)
	if err == nil {
		t.Fatal("expected error for expired token, got nil")
	}
}

func TestValidateJWT_InvalidUUIDInSubject(t *testing.T) {
	secret := "invalid-subject-secret"

	claims := jwt.RegisteredClaims{
		Issuer:  "chirpy",
		Subject: "not-a-uuid",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("failed to sign token: %v", err)
	}

	_, err = ValidateJWT(tokenString, secret)
	if err == nil {
		t.Fatal("expected error for invalid UUID subject, got nil")
	}
}

func TestValidateJWT_MalformedToken(t *testing.T) {
	secret := "malformed-secret"

	_, err := ValidateJWT("this-is-not-a-jwt", secret)
	if err == nil {
		t.Fatal("expected error for malformed token, got nil")
	}
}
