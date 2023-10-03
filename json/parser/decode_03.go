package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

// to figure out the usage of DisallowUnknownFields in decoder
func DecodeSample_03() {

	file, err := os.Open("sample_00.json")
	if err != nil {
		fmt.Printf("failed to load file %v", err)
		return
	}
	defer file.Close()

	var f flight

	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&f)
	if err != nil {
		fmt.Printf("failed to Decode json file %v\n", err)
		return
	}

	fmt.Printf("\n%+#v\n", f)

}

// compare DisallowUnknownFields of decoder with unmarshaler method
func UnMarshalSample_03() {

	file, err := os.ReadFile("sample_00.json")
	if err != nil {
		fmt.Printf("failed to read file %v", err)
		return
	}
	var f interface{}
	err = json.Unmarshal(file, &f)
	if err != nil {
		fmt.Printf("failed to Decode json file %v", err)
		return
	}
	// fmt.Printf("the fare is: %#v\n", *f.Fare)
	fmt.Printf("\n%#v\n", f)

}
