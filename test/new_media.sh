#!/bin/bash

curl -d '{
  "name": "Media Name",
  "domain": "example.com",
  "country": "TW",
  "company_id": 1,
  "source": "Wikipedia", 
  "note": "Note"
}' http://localhost:8080/media

curl -d '{
<<<<<<< HEAD
=======
  "id": 2,
>>>>>>> 08e915eeb40a16b228d4174c47b791a197055e2d
  "name": "Media Name LONGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG",
  "domain": "example.com",
  "country": "TW",
  "company_id": 2,
  "source": "Wikipedia AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", 
  "note": "Note"
}' http://localhost:8080/media