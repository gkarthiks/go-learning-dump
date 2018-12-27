package main

import (
    "database/sql"
    "fmt"
    "log"
	"time"
    _ "github.com/lib/pq"
)

func main() {
    // Connect to the "bank" database.
    db, err := sql.Open("postgres", "postgresql://root@localhost:26257?sslmode=disable")
    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }

    // Create the "accounts" table.
    if _, err := db.Exec(
        "CREATE TABLE IF NOT EXISTS test_sample (id INT PRIMARY KEY, balance INT, date DATE)"); err != nil {
        log.Fatal(err)
	}
	
	layout := "20060102"
	str := "20180901"
	resultantDate, _ := time.Parse(layout, str)

	if _, err := db.Exec(
		"INSERT INTO test_sample (id, balance, date) VALUES (1, 3000,$1)", resultantDate); err != nil {
		log.Fatal(err)
	}


	// Insert one by one
    // for i := 1; i < 10 ; i++ {
    // 	if _, err := db.Exec(
    // 		"INSERT INTO accounts (id, balance) VALUES ($1, 3000)", i); err != nil {
    // 		log.Fatal(err)
    // 	}
    // }

	// // Delete the data one by one
	// for i := 1; i < 1000000; i++ {
	// 	fmt.Println("Deleting this row now $1", i)
	// 	if _, err := db.Exec(
	// 		"delete from accounts where id = $1", i ); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// Delete all at once
	// if _, err := db.Exec(
	// 	"delete from accounts" ); err != nil {
	// 	log.Fatal(err)
	// }


    // Print out the balances.
	// rows, err := db.Query("SELECT id, balance FROM accounts")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // defer rows.Close()
    // fmt.Println("Initial balances:")
    // for rows.Next() {
    //     var id, balance int
    //     if err := rows.Scan(&id, &balance); err != nil {
    //         log.Fatal(err)
    //     }
    //     fmt.Printf("%d %d\n", id, balance)
	// }
	

	rows, err := db.Query("SELECT id, balance, date FROM test_sample")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    fmt.Println("Initial balances:")
    for rows.Next() {
		var id, balance int
		var date string
        if err := rows.Scan(&id, &balance, &date); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%d %d %s\n", id, balance, date)
	}
}
