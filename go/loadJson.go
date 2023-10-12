package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount uint
}

func main() {
	var contacts []contactInfo = loadJson[contactInfo]("./contacts.json")
	fmt.Printf("\n%+v", contacts)

}

func loadJson[T contactInfo | purchaseInfo](path string) []T {
	// you will need to specify the appropriate struct to unmarshall to
	data, _ := os.ReadFile(path)
	var loaded = []T{}
	err := json.Unmarshal(data, &loaded)
	if err != nil {
		println("Error occured Unmarshalling: ", err)
	}

	return loaded
}
