package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Config struct for database connection details
type Config struct {
	Database struct {
		User     string `json:"user"`
		DBName   string `json:"dbname"`
		Password string `json:"password"`
		SSLMode  string `json:"sslmode"`
	} `json:"database"`
}

// Pokemon struct to hold data from the database
type Pokemon struct {
	ID               int
	Number           string
	Name             string
	Type1            string
	Type2            sql.NullString // Type2 can be null
	ImageURL         string
	OptimalAttackers sql.NullString
	Attacks          sql.NullString
}

// Reads the configuration for database connection from a JSON file
func readConfig() (Config, error) {
	var config Config
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

// Main function to connect to the database, query, process, and update records
func main() {
	config, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s",
		config.Database.User, config.Database.DBName, config.Database.Password, config.Database.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, number, name, type1, type2, image_url FROM pokemon`)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Pokemon
		if err := rows.Scan(&p.ID, &p.Number, &p.Name, &p.Type1, &p.Type2, &p.ImageURL); err != nil {
			log.Fatal("Failed to scan row: ", err)
		}

		types := []string{p.Type1}
		if p.Type2.Valid {
			types = append(types, p.Type2.String)
		}

		effectiveAttacks, optimalAttackers := getEffectiveness(types)
		effectiveAttacksJSON, _ := json.Marshal(effectiveAttacks)
		optimalAttackersJSON, _ := json.Marshal(optimalAttackers)

		_, err := db.Exec(`UPDATE pokemon SET attacks = $1 WHERE id = $2`,
			effectiveAttacksJSON, p.ID)
		if err != nil {
			log.Fatal("Failed to update record: ", err)
		}

		_, err = db.Exec(`UPDATE pokemon SET optimalattackers = $1 WHERE id = $2`, optimalAttackersJSON, p.ID)
		if err != nil {
			log.Fatal("Failed to update record: ", err)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows: ", err)
	}
}
