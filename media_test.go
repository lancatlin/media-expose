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
	response := curl("POST", "/api/companies", company)
	assert.Equal(t, 200, response.StatusCode, resBody(response))

	media := `{
		"name": "Media Name",
		"domain": "example.com",
		"country": "TW",
		"company_id": 1,
		"source": "Wikipedia", 
		"note": "Note"
	}`
	response = curl("POST", "/api/media", media)
	assert.Equal(t, 200, response.StatusCode, resBody(response))

	response = curl("GET", "/api/media/1", "")
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

	response := curl("POST", "/api/media", media)
	assert.Error(t, ErrCompanyNotExist, resBody(response))
}
