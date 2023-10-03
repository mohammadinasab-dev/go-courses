package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
read multiple json object in the same stream.
decoder can decode this kind of data stream.
*/
func DecodeSample_02() {

	file, err := os.Open("sample_01.txt")
	if err != nil {
		fmt.Printf("failed to load file %v", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	for {
		var f flight

		err = decoder.Decode(&f)
		if err != nil {
			fmt.Printf("failed to Decode json file %v\n", err)
			return
		}

		fmt.Printf("\n%+#v\n", f)
	}

}

/*
read multiple json object in the same stream.
unmarshaler can not unmarshal this kind of data stream.
*/
func UnMarshalSample_02() {

	file, err := os.ReadFile("sample_01.txt")
	if err != nil {
		fmt.Printf("failed to read file %v", err)
		return
	}
	for {

		var f flight
		err = json.Unmarshal(file, &f)
		if err != nil {
			fmt.Printf("failed to Decode json file %v", err)
			return
		}
		fmt.Printf("\n%#v\n", f)
	}

}
