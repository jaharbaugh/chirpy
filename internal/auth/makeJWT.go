package auth

import(
	"time"
	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error){
	var claims jwt.RegisteredClaims
	claims.Issuer = "chirpy"
	claims.IssuedAt = jwt.NewNumericDate(time.Now().UTC())
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().UTC().Add(expiresIn))
	claims.Subject = userID.String()
	
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := newToken.SignedString([]byte(tokenSecret))
	if err != nil{
		return "", err
	}

	return signed, nil
}