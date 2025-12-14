package auth

import (
	"net/http"
	"errors"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error){
		headerString := headers.Get("Authorization")
		if headerString == ""{
			return "", errors.New("Header not set")
		}

		pref := "Bearer "
		tokenString := ""

		if strings.HasPrefix(headerString, pref){
			tokenString = strings.TrimPrefix(headerString, pref)
			tokenString = strings.TrimSpace(tokenString)
			if tokenString == ""{
				return "", errors.New("Missing token string")
			}
		}else{
			return "", errors.New("Incorrect prefix")
		}
		
	return tokenString, nil


}