package main

import "time"

type Company struct {
	ID              uint
	CreatedAt       time.Time
	Name            string
	Country         string
	Owner           string
	Shareholders    string
	InvestedByChina bool
}
