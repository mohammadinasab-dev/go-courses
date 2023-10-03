package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
	UNKNOWN INPUT

	- string or number
*/

func UnMarshalSample_0000() {

	file, err := os.ReadFile("sample_00.json")
	if err != nil {
		fmt.Printf("failed to read file %v", err)
		return
	}
	var f sample_01
	err = json.Unmarshal(file, &f)
	if err != nil {
		fmt.Printf("failed to Decode json file %v", err)
		return
	}
	// fmt.Printf("the fare is: %#v\n", *f.Fare)
	fmt.Printf("\n%#+v\n", f)

}
