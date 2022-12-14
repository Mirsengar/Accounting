package helpers

import (
	"Accounting/models"
	"encoding/json"
	"io/ioutil"
)

func SaveUser(user *models.User) {
	data, err := json.Marshal(user)
	if err != nil {
		logger.Println("User data is invalid. Unable to Marshal")
		logger.Fatalln(err)
	}
	err = ioutil.WriteFile("credentials.json", data, 0755)
	if err != nil {
		logger.Println("Unable to write user to credentials.json")
		logger.Fatalln(err)
	}
}

func GetUser() models.User {
	ba, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		logger.Println("Unable to read file: credentials.json")
		logger.Fatalln(err)
	}
	user := models.User{}
	err = json.Unmarshal(ba, &user)
	return user
}