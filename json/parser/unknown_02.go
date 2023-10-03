package parser

import (
	"encoding/json"
	"log"
)

/*
	The server sent two almost distinct messages, one for success and the other for failure.
	The exception is that both success and failure have a common field named 'status'.

*/


func (x *combinations) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &x.success); err != nil {
		return err
	} //else{
	// 	return nil
	// }

	if x.success.Status == "ok" { // we absoloutly should check the status 
		return nil
	}
	x.success = success{}
	return json.Unmarshal(b, &x.failure)
}


func Unmarshal_02() {
	r := make([]combinations, 0)
	if err := json.Unmarshal(data2, &r); err != nil {
		log.Println(err)
	}
	log.Printf("%+v", r)
}
