package db

import (
	"database/sql"
	"fmt"
	"log"

	// Justified.
	_ "github.com/lib/pq"
	psh "github.com/platformsh/config-reader-go/v2"
	libpq "github.com/platformsh/config-reader-go/v2/libpq"
)

func createConn() *sql.DB {
	config, err := psh.NewRuntimeConfig()
	checkErr(err)

	credentials, err := config.Credentials("postgresql")
	checkErr(err)

	// Retrieve the formatted credentials.
	formatted, err := libpq.FormattedCredentials(credentials)
	checkErr(err)
	db, err := sql.Open("postgres", formatted)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	return db
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Setup the database - Check if there are any tables and create them if not
func Setup() int {

	fmt.Println("Running db/Setup")
	// fmt.Println("Connecting to " + hostname + " on db " + database + ".")
	db := createConn()
	defer db.Close()
	var tableCount int
	err := db.QueryRow(`SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public'`).Scan(&tableCount)
	if err != nil {
		log.Fatal(err)
	}
	err = nil

	if tableCount == 0 {
		fmt.Println("Table count returned 0, creating tables.")
		_, err := db.Query(`
		DROP TABLE IF EXISTS posts;
		DROP TABLE IF EXISTS users;
		DROP TABLE IF EXISTS settings;

		
		
		-- CREATE TYPE post_status AS ENUM ('published', 'draft');
		
		CREATE TABLE posts 
		(
			id SERIAL PRIMARY KEY,
			author INT4,
			posted TIMESTAMPTZ,
			updated TIMESTAMPTZ,
			title TEXT,
			content TEXT,
			status post_status default 'draft',
			url TEXT,
			excerpt TEXT,
			meta XML
		);
		CREATE TABLE users
		(
			id SERIAL PRIMARY KEY,
			username VARCHAR UNIQUE,
			password VARCHAR,
			firstname VARCHAR,
			lastname VARCHAR,
			created TIMESTAMPTZ,
			updated TIMESTAMPTZ,
			last_login TIMESTAMPTZ,
			logins XML,
			settings XML
		);
		CREATE TABLE settings
		(
			id SERIAL PRIMARY KEY,
			setting VARCHAR,
			value TEXT
		)`)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Created database tables...")
	}
	fmt.Println("Tables exist, not creating anything.")
	return tableCount

}
