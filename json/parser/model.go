package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type sample_01 struct {
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
	lop json.RawMessage
}

type failure struct {
	Status  string
	ErrCode string
	Reason  string
}

type combinations struct {
	success
	failure
}

type response struct {
	result interface{}
}

// var data2 []byte = []byte(`[{"status":"ok","ReserveNumber":["125", "456"]}, {"status":"not ","ErrCode":"502", "Reason":"That's it"}]`)

var data2 []byte = []byte(`[{"status":"not ok","ErrCode":"502", "Reason":"null"}]`)

// var data3 []byte = []byte(`[{"status":"ok","ReserveNumber":["125", "456"]}]`)

// var data2 []byte = []byte(`[{"hello":"Bye"}]`)
