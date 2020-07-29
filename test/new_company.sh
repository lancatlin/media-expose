#!/bin/bash

curl -d '{
  "id": 2,
  "name": "Company Name",
  "country": "CN",
  "owner": "Owner",
  "shareholders": "Shareholders",
  "invested_by_china": true
}' http://localhost:8080/companies