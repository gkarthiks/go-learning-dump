package main

import (
    "database/sql"
    "log"
	_ "github.com/lib/pq"
	// "time"
	"fmt"
	// "strconv"
)

func main() {
    // Connect to the "bank" database.
    db, err := sql.Open("postgres", "postgresql://root@localhost:26257?sslmode=disable")
    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }

    // Create the "accounts" table.
    if _, err := db.Exec(
        "CREATE TABLE IF NOT EXISTS glassdoor (id INT PRIMARY KEY, ceoname STRING, featuredReviewoverall INT, name STRING, sectorName STRING, website STRING, date DATE, FILENAME STRING)"); err != nil {
        log.Fatal(err)
	}

	// firstRow := []string{"Doug McMillon", "4", "Walmart", "Retail", "www.walmart.com", "20180901", "glassdoor_enriched_20180901.csv"}
	// secondRow := []string{"Brian Cornell", "5", "Target", "None", "www.target.com", "20180901", "glassdoor_enriched_20180901.csv"}
	// thirdRow := []string{"Timothy J. Sloan", "5", "Wells Fargo", "None", "www.wellsfargo.com", "20180901", "glassdoor_enriched_20180901.csv"}
	// fourthRow := []string{"Robert", "5", "US Army", "None", "www.army.mil", "20180901", "glassdoor_enriched_20180901.csv"}
	// fifthRow := []string{"Jeff Bezos", "5", "Amazon", "Information Technology","www.amazon.jobs", "20180901", "glassdoor_enriched_20180901.csv"}
	// sixthRow := []string{"Randall L. Stephenson", "5", "AT&T", "Telecommunications", "www.att.jobs", "20180901", "glassdoor_enriched_20180901.csv"}
	// seventhRow := []string{"Craig Menear", "5", "The Home Depot", "None", "www.homedepot.com", "20180901", "glassdoor_enriched_20180901.csv"}
	// eigthRow := []string{"Brian T. Moynihan", "4", "Bank of America", "Finance", "www.bankofamerica.com","20180901", "glassdoor_enriched_20180901.csv"}
	// ninethRow := []string{"Virginia Rometty", "5", "IBM", "None", "www.ibm.com", "20180901", "glassdoor_enriched_20180901.csv"}
	// tenthRow := []string{"Kevin Johnson", "4", "Starbucks Restaurants", "Bars & Food Services", "www.starbucks.com", "20180901", "glassdoor_enriched_20180901.csv"}
	// tenSliceSlice := [][]string{firstRow, secondRow, thirdRow, fourthRow, fifthRow, sixthRow, seventhRow, eigthRow, ninethRow, tenthRow}

	// dateLayout := "20060102"


	// for i := 0; i < len(tenSliceSlice); i++ {
	// 	resultant, _ := time.Parse(dateLayout, tenSliceSlice[i][5])
	// 	review, _ := strconv.Atoi(tenSliceSlice[i][1])
	// 	if _, err := db.Exec(
	// 		"INSERT INTO glassdoor VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",i, tenSliceSlice[i][0], review, tenSliceSlice[i][2], tenSliceSlice[i][3], tenSliceSlice[i][4], resultant, tenSliceSlice[i][6]); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

    // Print out the balances.
    rows, err := db.Query("SELECT id, featuredReviewoverall, ceoname, name, sectorName, website, filename FROM glassdoor")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    fmt.Println("glassdoor data:")
    for rows.Next() {
		var id, featuredReviewoverall int 
		var ceoname, name, sectorName, website, filename string
		// , , , , , date

        if err := rows.Scan(&id, &featuredReviewoverall, &ceoname, &name, &sectorName, &website, &filename); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d Rating: %d CEO: %s \t\t  Name:%s \t\t Sector: %s %s\n", id, featuredReviewoverall, ceoname, name, sectorName, website)
	}
	
}
