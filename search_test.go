package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchByDomain(t *testing.T) {
	openTestDB()
	media := Media{
		Domain: "example.com",
	}
	assert.NoError(t, db.Create(&media).Error)
	domain := "example.com"
	res := curl("GET", fmt.Sprintf("/api/search?domain=%s", domain), "")
	assert.Equal(t, 200, res.StatusCode)

	var newMedia Media
	dec := json.NewDecoder(res.Body)
	dec.Decode(&newMedia)

	assert.Equal(t, media.Domain, newMedia.Domain)
}

func TestSearchNotFound(t *testing.T) {
	openTestDB()
	media := Media{
		Domain: "example.com",
	}
	assert.NoError(t, db.Create(&media).Error)
	domain := "another.com"
	response := curl("GET", "/api/search?domain="+domain, "")
	assert.Equal(t, 404, response.StatusCode, resBody(response))
}
