package main

import (
	"fmt"
	"encoding/json"
	"jaxf-github.fanatics.corp/TestBirthday/Model"

)


func main() {

	data := []byte(`{"ValuesNew" : [["Doe", "Joe", "1981/10/19"],["Some", "Body", "1953/09/29"]]}`)
	var pDetails Model.PersonDobDetails
	err := json.Unmarshal(data, &pDetails)
	if err != nil {
	fmt.Printf("Not working: %v\n", err)
	}

fmt.Println( findTodayPeopleBirthday(&pDetails))

}


