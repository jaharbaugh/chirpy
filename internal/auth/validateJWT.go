package auth

import(
	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error){
	var claims jwt.RegisteredClaims

	_, err := jwt.ParseWithClaims(
    	tokenString,
    	&claims,
    	func(t *jwt.Token) (interface{}, error) {
       	 return []byte(tokenSecret), nil
    	},
	)
	if err != nil{
		return uuid.Nil, err
	}
	
	parseID, err := uuid.Parse(claims.Subject)
	if err != nil{
		return uuid.Nil, err
	}

	return parseID, nil
}