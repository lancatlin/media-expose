#!/bin/bash

curl -d '{
  "id": 2,
  "name": "Media Name",
  "domain": "example.com",
  "country": "TW",
  "company_id": 2,
  "source": "Wikipedia", 
  "note": "Note"
}' http://localhost:8080/media