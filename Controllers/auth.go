package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Isauthenticated(w http.ResponseWriter, r *http.Request) {

	request, _ := http.NewRequest("POST", "https://svc-auth.rel.yuso.cloud/driver/auth", r.Body)
	request.Header.Set("content-Type", "application.json")
	Client := &http.Client{}
	response, err := Client.Do(request)

	if err != nil {

		fmt.Printf("the HTTP request failed with error %s\n", err)
	} else {

		data, _ := ioutil.ReadAll(response.Body)
		w.WriteHeader(response.StatusCode)
		w.Write([]byte(data))
	}

}

