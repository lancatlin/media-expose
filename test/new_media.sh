#!/bin/bash

curl -d '{
  "id": 2,
  "name": "Media Name",
  "domain": "example.com",
  "country": "TW",
  "company": {
    "id": 2,
    "name": "Company Name",
    "country": "CN",
    "owner": "Owner",
    "shareholders": "Shareholders",
    "invested_by_china": true
  },
  "source": "Wikipedia", 
  "note": "Note"
}' http://localhost:8080/media