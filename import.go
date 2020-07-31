package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func ImportCSVFromFile(filename string, processor func(func(string) string) error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	ImportCSV(file, processor)
	os.Exit(0)
}

func ImportCSV(r io.Reader, processor func(func(string) string) error) {
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

		if err = processor(col); err != nil {
			fmt.Println(err.Error(), row)
		}
	}
}

// ImportCompanies from CSV file
func ImportCompanies(col func(string) string) (err error) {
	company := Company{
		Meta: Meta{
			ID:      MustParseUInt(col("id")),
			Name:    col("name"),
			Country: col("country"),
			Source:  col("source"),
			Note:    col("note"),
		},
		Owner:           col("owner"),
		Shareholders:    col("shareholders"),
		InvestedByChina: MustParseBool(col("invested_by_china")),
	}

	if err = company.Verify(); err != nil {
		return
	}
	err = db.Create(&company).Error
	if err != nil {
		return
	}
	fmt.Printf("success import %s\n", company.Name)
	return nil
}

func ImportMedia(col func(string) string) error {
	media := Media{
		Meta: Meta{
			Name:    col("name"),
			Country: col("country"),
			Source:  col("source"),
			Note:    col("note"),
		},
		Domain:    col("domain"),
		CompanyID: MustParseUInt(col("company_id")),
	}
	if err := media.Verify(); err != nil {
		return err
	}

	if err := db.Create(&media).Error; err != nil {
		return err
	}

	fmt.Printf("success import %s\n", media.Name)
	return nil
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
