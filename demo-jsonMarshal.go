package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {

	data, _ := json.Marshal(&employee{101, "witty king", "01100000", "witty@mail.com"})
	fmt.Println(string(data))
}
