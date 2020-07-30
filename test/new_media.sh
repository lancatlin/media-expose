#!/bin/bash

curl -d '{
  "name": "Media Name",
  "domain": "example.com",
  "country": "TW",
  "company_id": 1,
  "source": "Wikipedia", 
  "note": "Note"
}' http://localhost:8080/api/media

curl -d '{
  "name": "Media Name LONGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG",
  "domain": "example.com",
  "country": "TW",
  "company_id": 2,
  "source": "Wikipedia AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", 
  "note": "Note"
}' http://localhost:8080/api/media