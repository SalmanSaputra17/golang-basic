package main

import(
	"encoding/json"
	"net/http"
	"net/url"
	"bytes"
	"fmt"
)

var baseUrl = "http://localhost:8080"

type student struct {
	ID    string
    Name  string
    Grade int
}

func fetchUsers() ([]student, error) {
	var err error
	var client = &http.Client{}
	var data []student

	request, err := http.NewRequest("POST", baseUrl + "/users", nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchUserByID(ID string) (student, error) {
	var err error
	var client = &http.Client{}
	var data student

	var param = url.Values{}
	param.Set("ID", ID)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseUrl + "/user", payload)

	if err != nil {
		return data, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)

	if err != nil {
		return data, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	// fetch all users
	var users, err_1 = fetchUsers()

	if err_1 != nil {
		fmt.Println("Error!", err_1.Error())
        return
	}

	for _, each := range users {
        fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
    }

    // fetch spesific user by ID
    var user, err_2 = fetchUserByID("E001")

    if err_2 != nil {
        fmt.Println("Error!", err_2.Error())
        return
    }

    fmt.Printf("\nID: %s\t Name: %s\t Grade: %d\n", user.ID, user.Name, user.Grade)
}