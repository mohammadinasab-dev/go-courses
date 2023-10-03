package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

/*
here we can modify the struct with different json tag and options.
*/
type flight struct {
	ID          string `json:"-"`
	Origin      string `json:"origin,omitempty"`
	Destination string `json:"-"`
	Number      string
	Class       SliceOrString
	// Class    string
	Airline  string
	AirCraft string      `json:",omitempty"`
	Seats    IntOrString //`json:"-"`
	// Seats    string
	// Fare     Fare   `json:"fare,omitempty"`
	Fare *Fare `json:"fare,omitempty"`
	// Fare          []Fare
	Tax           Tax //`json:"-"`
	TotalPrice    IntOrString
	DepartureDate string
	ArrivalDate   string
	CreatedAt     customTime
	DeletedAt     time.Time
}

type Fare struct {
	Adl int64
	Chd int64
	Inf int64
}

type Tax struct {
	Adl int64
	Chd int64
	Inf int64
}

type customTime struct {
	time.Time
}

func (x customTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", x.Format("2006-01-02"))), nil
}

func (x *customTime) UnmarshalJSON(b []byte) error {
	var err error
	x.Time, err = time.Parse("2006-01-02", strings.Trim(string(b), "\""))
	return err
}

type IntOrString int

func (x *IntOrString) UnmarshalJSON(b []byte) error {
	var v int
	err := json.Unmarshal(bytes.Trim(b, `"`), &v)
	*x = IntOrString(v)
	return err
}

type SliceOrString []string

func (x *SliceOrString) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		var v string
		err := json.Unmarshal(b, &v)
		*x = SliceOrString{v}
		return err
	}
	var v []string
	err := json.Unmarshal(b, &v)
	*x = SliceOrString(v)
	return err
}

type success struct {
	Status        string
	ReserveNumber []string
}

type failure struct {
	Status  string
	ErrCode string
	Reason  string
}

type response struct {
	result interface{}
}

func (x *response) UnmarshalJSON(b []byte) error {
	var success success
	if err := json.Unmarshal(b, &success); err != nil {
		return err
	}

	//Think about it why these if condition is `necessary`!
	if success.Status == "ok" {
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
