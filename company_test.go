package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCompany(t *testing.T) {
	openTestDB()
	body := `{
		"id": 1493,
		"name": "Company Name",
		"country": "CN",
		"owner": "Owner",
		"shareholders": "Shareholders",
		"invested_by_china": true
	}`
	response := curl("POST", "/api/companies", body)
	data := resBody(response)
	assert.Equal(t, 200, response.StatusCode, data)
	assert.Equal(t, "1493", data)

	response = curl("GET", "/api/companies/1493", "")
	assert.Equal(t, 200, response.StatusCode)
}

func TestDupCompany(t *testing.T) {
	openTestDB()

	body := `{
		"id": 163032,
		"name": "A company",
		"country": "CN",
		"owner": "Owner",
		"shareholders": "Shareholders",
		"invested_by_china": true
	}`
	response := curl("POST", "/api/companies", body)
	assert.Equal(t, 200, response.StatusCode, resBody(response))

	response = curl("POST", "/api/companies", body)
	assert.Equal(t, 403, response.StatusCode, resBody(response))
}
