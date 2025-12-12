package auth

import(
	"github.com/alexedwards/argon2id"
	"log"
)

func CheckPasswordHash(password string, hash string) (bool, error){
	
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		log.Fatal(err)
		return match, err
	}

	log.Printf("Match: %v", match)
	return match, nil

}