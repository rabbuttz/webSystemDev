package main

import (
	"fmt"

	"getTwitterTrend"
)

func main() {
	trnd := getTwitterTrend.GetTrends(23424856)
	fmt.Println(trnd.Trends)
}