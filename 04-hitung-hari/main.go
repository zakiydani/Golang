package main

import (
	"fmt"
	"time"
)

func main() {
	March := time.Date(2020, time.March, 3, 4, 20, 0, 0, time.UTC)
	June := time.Date(2020, time.June, 10, 13, 13, 0, 0, time.UTC)

	Distance := June.Sub(March)

	Days := int(Distance.Hours() / 24)

	fmt.Printf("3 Maret 2020 - 10 Juni 2020 ada %d jumlah hari\n", Days)
}