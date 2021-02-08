package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/theovidal/bacbot/lib"
)

type Response struct {
	Errors  []interface{}
	Message string
	Token   string
	Data    Data
}

type Data struct {
	Homeworks []Homework
	Timetable []Lesson
	Contents  Contents
}

func MakeRequest(query string) (result Response, err error) {
	request, _ := http.NewRequest(
		"POST",
		os.Getenv("PRONOTE_API")+"/graphql",
		strings.NewReader(query),
	)

	request.Header.Add("Content-Type", "application/json")

	var response *http.Response
	for {
		var currentResult Response
		token := lib.Cache.Get(context.Background(), "token").Val()
		request.Header.Set("Token", token)

		response, err = lib.DoRequest(request)
		if err != nil {
			return
		}
		var bytes []byte
		bytes, _ = ioutil.ReadAll(response.Body)
		response.Body.Close()

		_ = json.Unmarshal(bytes, &currentResult)

		fmt.Println(response.Request.Host, response.StatusCode, string(bytes))

		fmt.Println(len(currentResult.Errors), currentResult.Message)
		if response.StatusCode == 200 && len(currentResult.Errors) == 0 && currentResult.Message == "" {
			result = currentResult
			break
		}

		err = Login()
		if err != nil {
			return
		}
	}

	return
}

func Login() error {
	query, _ := json.Marshal(map[string]string{
		"url":      os.Getenv("PRONOTE_SERVER"),
		"cas":      os.Getenv("PRONOTE_CAS"),
		"username": os.Getenv("PRONOTE_USER"),
		"password": os.Getenv("PRONOTE_PASSWORD"),
	})

	request, _ := http.NewRequest(
		"POST",
		os.Getenv("PRONOTE_API")+"/auth/login",
		bytes.NewReader(query),
	)

	request.Header.Add("Content-Type", "application/json")
	response, err := lib.DoRequest(request)
	if err != nil {
		return err
	}

	var bytes []byte
	bytes, _ = ioutil.ReadAll(response.Body)
	response.Body.Close()

	var result Response
	_ = json.Unmarshal(bytes, &result)

	lib.Cache.Set(context.Background(), "token", result.Token, 0)

	return nil
}
