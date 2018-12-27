package main

import (
	"fmt"
	"time"
)

func main() {
	firstRow := []string{"Doug McMillon", "4", "Walmart", "Retail", "www.walmart.com", "20180901", "glassdoor_enriched_20180901.csv"}
	secondRow := []string{"Brian Cornell", "5", "Target", "None", "www.target.com", "20180901", "glassdoor_enriched_20180901.csv"}
	thirdRow := []string{"Timothy J. Sloan", "5", "Wells Fargo", "None", "www.wellsfargo.com", "20180901", "glassdoor_enriched_20180901.csv"}
	fourthRow := []string{"Robert", "5", "US Army", "None", "www.army.mil", "20180901", "glassdoor_enriched_20180901.csv"}
	fifthRow := []string{"Jeff Bezos", "5", "Amazon", "Information Technology","www.amazon.jobs", "20180901", "glassdoor_enriched_20180901.csv"}
	sixthRow := []string{"Randall L. Stephenson", "5", "AT&T", "Telecommunications", "www.att.jobs", "20180901", "glassdoor_enriched_20180901.csv"}
	seventhRow := []string{"Craig Menear", "5", "The Home Depot", "None", "www.homedepot.com", "20180901", "glassdoor_enriched_20180901.csv"}
	eigthRow := []string{"Brian T. Moynihan", "4", "Bank of America", "Finance", "www.bankofamerica.com","20180901", "glassdoor_enriched_20180901.csv"}
	ninethRow := []string{"Virginia Rometty", "5", "IBM", "None", "www.ibm.com", "20180901", "glassdoor_enriched_20180901.csv"}
	tenthRow := []string{"Kevin Johnson", "4", "Starbucks Restaurants", "Bars & Food Services", "www.starbucks.com", "20180901", "glassdoor_enriched_20180901.csv"}

	tenSliceSlice := [][]string{firstRow, secondRow, thirdRow, fourthRow, fifthRow, sixthRow, seventhRow, eigthRow, ninethRow, tenthRow}
	layout := "20060102"

	for i := 0; i < len(tenSliceSlice); i++ {
		resultant, _ := time.Parse(layout, tenSliceSlice[i][5])
		fmt.Println("INSERT INTO test_sample (id, ceoname, featuredReviewoverall, name, sectorName, website , date, FILENAME) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",i, tenSliceSlice[i][0], tenSliceSlice[i][1], tenSliceSlice[i][2], tenSliceSlice[i][3], tenSliceSlice[i][4], resultant, tenSliceSlice[i][6])
	}
	
}