package parser

import (
	"encoding/json"
	"log"
)

/*
	The server sent two almost distinct messages, one for success and the other for failure.
	The exception is that both success and failure have a common field named 'status'.

*/

func (x *response) UnmarshalJSON(b []byte) error {
	var success success
	if err := json.Unmarshal(b, &success); err != nil {
		return err
	}

	if success.Status == "ok" { // we absoloutly should check the status
		x.result = success
		return nil
	}
	var failure failure
	if err := json.Unmarshal(b, &failure); err != nil {
		return err
	}
	x.result = failure
	return nil
}

func Unmarshal_022() {
	r := make([]response, 0)
	if err := json.Unmarshal(data2, &r); err != nil {
		log.Println(err)
	}
	log.Printf("%+v", r)
}
