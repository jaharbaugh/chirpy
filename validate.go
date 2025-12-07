package main

import(
	"net/http"
	"encoding/json"
	"log"
)

func handlerValidate (w http.ResponseWriter, req *http.Request){ 
	if req.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	max_chars := 140

	type chirp struct {
        Body string `json:"body"`
    }

	decoder := json.NewDecoder(req.Body)
	chirpBody := chirp{}
	err := decoder.Decode(&chirpBody)
	if err != nil{
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	type responseInvalid struct {
		Error string `json:"error"`	
	}
	type responseValid struct{
		Valid bool   `json:"valid"`
	}
	
	var status int
	var res interface{}

	if len(chirpBody.Body) > max_chars{
		res = responseInvalid{
			Error: "Chirp is too long",
		}
		status = http.StatusBadRequest
	}else{
		res = responseValid{
			Valid: true,
		}
		status = http.StatusOK
	}

	dat, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}



	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
	w.Write(dat)

}