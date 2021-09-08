package storage

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// SecretHandler manage the GET and POST requests
func (ids Storage) SecretHandler(w http.ResponseWriter, r *http.Request) {
	var status int
	var ret string
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path == "/" {
			log.Println("Bad Request")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Bad, Bad Request!\n")
			return
		}
		ret, status = ids.secretGet(r.URL.Path[1:])
	case http.MethodPost:
		var bodyBytes []byte
		var err error
		if r.Body != nil {
			bodyBytes, err = ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("Body reading error: %v", err)
				return
			}
			defer r.Body.Close()
		}

		ret, status = ids.secretPost(bodyBytes)

	default:
		status = http.StatusMethodNotAllowed
		ret = "Method not allowed!\n"
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintln(w, ret)
}
