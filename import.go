package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
)

// ImportCompanies from CSV file
func ImportCompanies(r io.Reader) {
	dec := csv.NewReader(r)
	columnNames, err := dec.Read()
	if err != nil {
		log.Fatal(err)
	}
	columns := make(map[string]int)
	for i, col := range columnNames {
		columns[col] = i
	}

	for {
		row, err := dec.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		col := func(key string) string {
			return row[columns[key]]
		}

		company := Company{
			Meta: Meta{
				ID:     MustParseUInt(col("id")),
				Name:   col("name"),
				Source: col("source"),
				Note:   col("note"),
			},
			Owner:           col("owner"),
			Shareholders:    col("shareholders"),
			InvestedByChina: MustParseBool(col("invested_by_china")),
		}
		if err = company.Verify(); err != nil {
			fmt.Println(err.Error(), company)
			continue
		}
		err = db.Create(&company).Error
		if err != nil {
			fmt.Println(err.Error(), company)
			continue
		}
		fmt.Printf("success import %s\n", company.Name)
	}
}

func MustParseUInt(s string) uint {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return uint(n)
}

func MustParseBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	return b
}
