package main

import (
	"encoding/json"
	"log"
)

type People struct {
	Id   int64    `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags,omitempty"`
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {

	pBytes := []byte(`{
		"id":1,
		"name":"tabalt",
		"tags":[
			"Gopher",
			"Chinese"
		]
	}`)

	p := People{}
	if err := json.Unmarshal(pBytes, &p); err != nil {
		log.Fatal(err)
	}

	log.Printf("people %v", p)

	pbs, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("json bytes %s", pbs)
}
