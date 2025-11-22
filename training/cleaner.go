package main

import "encoding/json"


type product struct {
	Name    string `json:"name"`
	Price   int    `json:"price_cents"`
	InStock bool   `json:"in_stock"`
}

func cleanProductJSON(input string) string {
	var p product
	_ = json.Unmarshal([]byte(input), &p)
	m, _ := json.Marshal(p)
	return string(m)
}
