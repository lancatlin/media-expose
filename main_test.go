package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

func openTestDB() {
	conf := Configure{
		Mode: Memory,
	}
	OpenDB(conf)
}

func curl(method, target, body string) *http.Response {
	req := httptest.NewRequest(method, target, strings.NewReader(body))
	rec := httptest.NewRecorder()
	NewRouter().ServeHTTP(rec, req)
	return rec.Result()
}

func resBody(res *http.Response) string {
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}
