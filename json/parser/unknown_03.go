package parser

import (
	"encoding/json"
	"log"
)

/*
	The server sent two almost distinct messages, one for success and the other for failure.
	The exception is that both success and failure have a common field named 'status'.

*/


/*
	Here is the sample data stream. Try both one after the other.
*/
// var data2 []byte = []byte(`[{"status":"ok","ReserveNumber":["125", "456"]}, {"status":"not ","ErrCode":"502", "Reason":"That's it"}]`)
var data2 []byte = []byte(`[{"status":"not ok","ErrCode":"502", "Reason":"null"}]`)
// var data3 []byte = []byte(`[{"status":"ok","ReserveNumber":["125", "456"]}]`)
// var data2 []byte = []byte(`[{"hello":"Bye"}]`)



func UnMarshalUnKnown_03() {
	r := make([]response, 0)
	if err := json.Unmarshal(data2, &r); err != nil {
		log.Println(err)
	}
	log.Printf("%+v", r)
}
