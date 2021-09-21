package gateways

import (
	"encoding/json"
	"fmt"
	"homework/internal/app/models"
	"io/ioutil"
	"os"
)

func GetJson() models.User {
	var a models.User

	file, err := os.Open("data1.json")
	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &a)
	if err != nil {
		fmt.Println(err)
	}

	return a
}
