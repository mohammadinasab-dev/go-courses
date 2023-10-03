package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
	UNKNOWN INPUT

	- string or number
	- slice or string

	- you can modify the flight model to have field with custom data type (IntOrString, SliceOrString)
	- as already know these custom type implement the `Unmarshaler interface` ...
*/

func UnMarshalUnKnown_01() {

	file, err := os.ReadFile("sample_00.json")
	if err != nil {
		fmt.Printf("failed to read file %v", err)
		return
	}
	var f flight
	err = json.Unmarshal(file, &f)
	if err != nil {
		fmt.Printf("failed to Decode json file %v", err)
		return
	}
	// fmt.Printf("the fare is: %#v\n", *f.Fare)
	fmt.Printf("\n%#+v\n", f)

}
