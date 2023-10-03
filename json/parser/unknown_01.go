package parser

import (
	"encoding/json"
	"log"
)

// The server sent two distinct messages, one for success and the other for failure.

type success01 struct {
	ReserveNumber []string
}

type errorResult struct {
	ErrCode string
	Reason  string
}

type res struct {
	success01
	errorResult
}

var data []byte = []byte(`[{"ErrCode":"502", "Reason":"that's it"}]`)
// var data []byte = []byte(`[{"ReserveNumber":["125", "456"]}]`)

func Unmarshal_01() {
	r := make([]res, 0)
	if err := json.Unmarshal(data, &r); err != nil {
		log.Println(err)
	}
	log.Printf("%+v", r)
}
