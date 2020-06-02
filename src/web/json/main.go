package main

import(
	"fmt"
	"encoding/json"
)

type User struct {
	FullName string `json:"Name"`
	Age int
}

func jsonToStruct(s string) {
	var jsonData = []byte(s)
	var data User

	var result = json.Unmarshal(jsonData, &data)

	if result != nil {
		fmt.Println(result.Error())
		return
	}

	fmt.Println("user :", data.FullName)
    fmt.Println("age  :", data.Age)
}

func jsonToMapInterface(s string) {
	var jsonData = []byte(s)
	var data map[string]interface{}

	json.Unmarshal(jsonData, &data)

	fmt.Println("user :", data["Name"])
	fmt.Println("age  :", data["Age"])
}

func jsonToInterface(s string) {
	var jsonData = []byte(s)
	var data interface{}

	json.Unmarshal(jsonData, &data)

	var result = data.(map[string]interface{})	

	fmt.Println("user :", result["Name"])
	fmt.Println("age  :", result["Age"])
}

func jsonToArray(s string) {
	var jsonData = []byte(s)
	var data []User

	var result = json.Unmarshal(jsonData, &data)

	if result != nil {
		fmt.Println(result.Error())
		return
	}

	fmt.Println("user 1:", data[0].FullName)
	fmt.Println("user 2:", data[1].FullName)
}

func encodeToJson(data []User) string {
	var jsonData, err = json.Marshal(data)

	if err != nil {
		return err.Error()
	}

	return string(jsonData)
}

func main() {
	var jsonString = `{"Name": "Salman Saputra", "Age": 18}`
	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}

	jsonToStruct(jsonString)
	jsonToMapInterface(jsonString)
	jsonToInterface(jsonString)
	jsonToArray(encodeToJson(object))
}