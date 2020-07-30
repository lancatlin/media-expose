package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomain(t *testing.T) {
	m := Media{
		Domain: "example.com",
	}
	assert.False(t, m.domainNotExist())
	m.Domain = "notexistd.omain"
	assert.True(t, m.domainNotExist())
}

func TestNewMedia(t *testing.T) {
	openTestDB()
	company := `{
		"name": "Company Name",
		"country": "CN",
		"owner": "Owner",
		"shareholders": "Shareholders",
		"invested_by_china": true
	}`
	response := curl("POST", "/companies", company)
	assert.Equal(t, 200, response.StatusCode, resBody(response))

	media := `{
		"name": "Media Name",
		"domain": "example.com",
		"country": "TW",
		"company_id": 1,
		"source": "Wikipedia", 
		"note": "Note"
	}`
	response = curl("POST", "/media", media)
	assert.Equal(t, 200, response.StatusCode, resBody(response))
}

func TestCompanyNotFound(t *testing.T) {
	openTestDB()
	media := `{
		"name": "Media Name",
		"domain": "example.com",
		"country": "TW",
		"company_id": 1,
		"source": "Wikipedia", 
		"note": "Note"
	}`

	response := curl("POST", "/media", media)
	assert.Error(t, ErrCompanyNotExist, resBody(response))
}
