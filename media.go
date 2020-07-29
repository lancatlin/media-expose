package main

type Media struct {
	Meta

	Domain  string
	Country string

	Company   Company
	CompanyID uint `json:"company_id"`
}
