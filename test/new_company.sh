#!/bin/bash

curl -d '{
  "name": "Company Name",
  "country": "CN",
  "owner": "Owner",
  "shareholders": "Shareholders",
  "invested_by_china": true
}' http://localhost:8080/api/companies

curl -d '{
  "name": "Company Name",
  "country": "CN",
  "owner": "Owner",
  "shareholders": "Shareholders",
  "invested_by_china": true
}' http://localhost:8080/api/companies

curl -d '{
  "name": "Company Name TOO LONGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG",
  "country": "CN",
  "owner": "Owner",
  "shareholders": "Shareholders",
  "invested_by_china": true
}' http://localhost:8080/api/companies

curl -d '{
  "country": "CN",
  "owner": "Owner",
  "shareholders": "Shareholders",
  "invested_by_china": true
}' http://localhost:8080/api/companies