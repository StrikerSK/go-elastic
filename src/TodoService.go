package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func createData(object CustomInterface) string {
	url := HOST_URL + "/" + TODOS_INDEX + "/_doc"
	method := "POST"

	marshalledObject, _ := object.MarshalItem()
	payload := strings.NewReader(string(marshalledObject))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "nil"
	}
	log.Print("Custom object created successfully!")

	m := make(map[string]string)
	_ = json.Unmarshal(body, &m)

	fmt.Println(string(body))
	return m["_id"]
}

func getTodo(todoID string) (todo Todo) {

	url := HOST_URL + "/" + TODOS_INDEX + "/_doc/" + todoID
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if res.StatusCode == http.StatusOK {
		m := make(map[string]interface{})
		_ = json.Unmarshal(body, &m)

		todo.ResolveMap(m["_source"].(map[string]interface{}))
		return
	} else {
		log.Println("Something went wrong at getTodo()")
		return
	}
}
