package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Request struct {
	RegistrationIds []string `json:"registration_ids"`
}

type Response struct {
	Multicast_Id int `json:"multicast_id"`
	Success      int `json:"success"`
	Failure      int `json:"failure"`
}

func main() {

	fmt.Println("running on 8080")
	http.HandleFunc("/", HandleRequest)
	http.ListenAndServe(":8080", nil)

}

func HandleRequest(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	json_data := Request{}
	json.Unmarshal([]byte(body), &json_data)

	response := Response{108, 4, 96}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("request received")

	// introduce some latency for good measure
	time.Sleep(1250 * time.Millisecond)

	w.Write(b)

}
