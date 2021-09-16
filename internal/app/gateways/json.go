package gateways

import (
	"encoding/json"
	"fmt"
	"homework/internal/app/models"
	"io/ioutil"
	"os"
)

func GetJson(a *models.AddUser) (int, string, string, int, string) {
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
	return a.Id, a.FirstName, a.LastName, a.Age, a.Status
}
