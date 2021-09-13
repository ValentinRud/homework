package gateways

import (
	"encoding/json"
	"fmt"
	"homework/internal/app/models"
	"io/ioutil"
	"os"
)

var M *models.AddUser

func GetJson() {
	file, err := os.Open("data1.json")
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &M)
	if err != nil {
		fmt.Println(err)
	}
}