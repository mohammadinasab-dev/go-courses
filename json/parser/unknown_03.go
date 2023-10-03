package parser

import (
	"encoding/json"
	"fmt"
)

type raw []byte

func (x *raw) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &x)
	return err
}

type Person struct {
	Name string          `json:"name"`
	Age  int             `json:"age"`
	Data json.RawMessage `json:"data"`
}

func Raw() {
	jsonStr := `{
        "name": "John",
        "age": 30,
        "data": {
            "email": "john@example.com",
            "phone": "123-456-7890",
			"number": 123
        }
    }`

	var p Person
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p.Name) // Output: John
	fmt.Println(p.Age)  // Output: 30

	// Unmarshal the raw message into a map[string]string
	var data map[string]interface{}
	err = json.Unmarshal(p.Data, &data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data["email"]) // Output: john@example.com
	fmt.Println(data["phone"]) // Output: 123-456-7890
	fmt.Printf("%+#v\n", data["number"])
}
