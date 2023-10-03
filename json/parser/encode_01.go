package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

var s sample_01 = sample_01{
	ID: uuid.NewString(),
	// Origin:      "KIH",
	Destination: "THR",
	Number:      "20598",
	// Class:       "[EEC]",
	// Airline:  "زاگرس",
	// AirCraft: "Boeing 737",
	// Seats:    "5",
	// Fare: &Fare{
	// 	Adl: 12000000,
	// 	Chd: 9000000,
	// 	Inf: 4500000,
	// },
	Tax: Tax{
		Adl: 1200000,
		Chd: 900000,
		Inf: 30000,
	},
	DepartureDate: time.Now().String(),
	ArrivalDate:   time.Now().AddDate(0, 0, 12).GoString(),
	CreatedAt:     customTime{time.Now()},
	DeletedAt:     time.Now(),
}

func Encode_01() {

	err := json.NewEncoder(os.Stdout).Encode(s)
	if err != nil {
		log.Fatal("failed to marshal : ", err)
	}

}

func Marshal_01() {

	// sm , err := json.Marshal(s)
	sm, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		log.Fatal("failed to marshal : ", err)
	}
	fmt.Printf("json output is: %v\n", string(sm))
}
