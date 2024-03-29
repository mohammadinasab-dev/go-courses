package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
  - un-exported fields will be omitted during encoding and decoding in JSON. DONE

  - it is not necessary to define all the JSON keys in the struct. DONE
*/

/*
   JSON TAG CONSIDERATIONS

   - default tag is as the same as field name. DONE

   - in go, json tag is case-insensitive, json. DONE
     e.g. `json:"Origin"` - `json:"origin"` - `json:"oRigin"` are all equal

   - in go, the priority of JSON tag is higher than the field name itself. DONE

   - using the same json tag for different field cause the data lost. DONE

   - hyphen(`-`) as json tag tells the json package to ignore this specific field to decode and encoding. DONE

   - `omitempty` as json tag tells the json package to omit the field from json output if it has zero value.

   - `string` tells the json package to encode the field to the corresponding data type of just int, float or boolean. DONE

   - `,` is the separator character in json tag. DONE
*/

/*
   - decoder use io.reader as an input argument in which unmarshal use []byte. DONE

   - decoder can read multiple json object in which unmarshal can't. DONE

   - decoder give us more option and flexibility. DONE
     e.g. DisallowUnknownFields
*/

// simple decoder without any specific options
func DecodeSample_01() {

	file, err := os.Open("sample_00.json")
	if err != nil {
		fmt.Printf("failed to load file %v", err)
		return
	}
	defer file.Close()

	var f flight
	if err := json.NewDecoder(file).Decode(&f); err != nil {
		fmt.Printf("failed to Decode json file %v\n", err)
		return
	}
	// if err != nil {
	// 	fmt.Printf("failed to Decode json file %v\n", err)
	// 	return
	// }
	fmt.Printf("\n%+#v\n", f)

}

// simple unmarshaler without any specific options
func UnMarshalSample_01() {

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
	fmt.Printf("\n%#v\n", f)

}
