package storage

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Storage struct hold the map and the pointer to the file
type Storage struct {
	IDs map[string]string
	fn  string
}

// InitMapFile load the map stored on file, or
// it creates the empty file
func (ids *Storage) InitMapFile(fn string) error {
	var fp *os.File
	fn = filepath.Clean(fn)

	s, err := os.Stat(fn)
	defer fp.Close()

	if os.IsNotExist(err) {
		log.Printf("%s does not exists, creating it...\n", fn)
		fp, err = os.Create(fn)
		if err != nil {
			log.Panic(err)
		}
	} else {
		log.Printf("Found %s, proceed", fn)
		fp, err = os.Open(fn)
		if err != nil {
			log.Panic(err)
		}
	}

	if s.Size() > 0 {
		buf := make([]byte, s.Size())
		_, err := io.ReadFull(fp, buf)
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
			log.Panicf("Panic trying to read the content of %s: %v\n", fn, err)
		}
		err = json.Unmarshal(buf, &ids.IDs)
		if err != nil {
			log.Panic(err)
		}
	}

	log.Printf("Map Loaded")
	ids.fn = fn
	return nil
}

// UpdateMap save the map with the IDs on the file
func (ids *Storage) updateMap(key, action string) (string, error) {
	var id string
	f, err := os.OpenFile(ids.fn, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		log.Printf("Error opening the file %s: %v\n", ids.fn, err)
		return "", err
	}
	defer f.Close()
	switch action {
	case "del":
		delete(ids.IDs, key)
	case "put":
		id = fmt.Sprintf("%x", md5.Sum([]byte(key)))
		ids.IDs[id] = key
	}
	tmp, err := json.Marshal(ids.IDs)
	if err != nil {
		log.Printf("Error Marshalling the map to json %s: %v\n", ids.IDs, err)
		return "", err
	}

	n, err := f.Write(tmp)
	if err != nil {
		log.Printf("Error writing in the file %s: %v\n", ids.fn, err)
		return "", err
	}

	log.Printf("Updated file %v, written %d bytes, err: %v\n", ids.fn, n, err)
	return id, nil
}

func (ids Storage) secretGet(id string) (string, int) {
	v, found := ids.IDs[id]
	if found {
		log.Printf("Found Key searched: %s\n", v)
		_, err := ids.updateMap(id, "del")
		if err != nil {
			log.Printf("Error deleting key from json: %v\n", err)
		}
		return fmt.Sprintf("{\"data\": \"%s\"}", v), http.StatusOK
	}
	log.Printf("Not found key searched: %s, map: %v\n", id, ids.IDs)

	return enc(""), http.StatusNotFound
}

func (ids Storage) secretPost(bodyBytes []byte) (string, int) {
	var secret map[string]string

	err := json.Unmarshal(bodyBytes, &secret)
	if err != nil {
		log.Printf("Error unmarshalling the body %s, %v\n", bodyBytes, err)
	}

	v := secret["plain_text"]
	if v == "" {
		log.Printf("Missing expected key (plain_text) in POST Request!!!\n")
		return "Unknown Key", http.StatusInternalServerError
	}
	log.Printf("Found key: %v\n", secret)
	id, err := ids.updateMap(v, "put")
	if err != nil {
		log.Printf("Error adding key to json: %v\n", err)
	}
	return enc(id), http.StatusOK
}

func enc(v string) string {
	ret, err := json.Marshal(map[string]string{"data": v})
	if err != nil {
		return ""
	}
	return string(ret)
}
