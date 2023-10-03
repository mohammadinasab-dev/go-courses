package parser

import (
	"encoding/json"
	"log"
)

/*
	- The server sent two distinct messages, one for reserve and the other for failure.
	- Think about it! here in this file there is no need to implement th `Unmarshaler interface`.
*/ 

type reserve struct {
	ReserveNumber []string
}

type errorResult struct {
	ErrCode string
	Reason  string
}

type res struct {
	reserve
	errorResult
}

// Here is the sample data stream. Try both one after the other.
var data []byte = []byte(`[{"ErrCode":"502", "Reason":"that's it"}]`)
// var data []byte = []byte(`[{"ReserveNumber":["125", "456"]}]`)

func UnMarshalUnKnown_02() {
	r := make([]res, 0)
	if err := json.Unmarshal(data, &r); err != nil {
		log.Println(err)
	}
	log.Printf("%+v", r)
}
