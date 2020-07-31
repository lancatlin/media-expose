package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCompany(t *testing.T) {
	openTestDB()
	body := `{
		"name": "Company Name",
		"country": "CN",
		"owner": "Owner",
		"shareholders": "Shareholders",
		"invested_by_china": true
	}`
	response := curl("POST", "/api/companies", body)
	data := resBody(response)
	assert.Equal(t, 200, response.StatusCode, data)
	assert.Equal(t, "1", data)

	response = curl("GET", "/api/companies/1", "")
	assert.Equal(t, 200, response.StatusCode)
}

func TestDupCompany(t *testing.T) {
	openTestDB()

	body := `{
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
