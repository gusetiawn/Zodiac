package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
	"time"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "marvel" //dbname existing di localku
)

var db *sql.DB

type PageData struct {
	Name      string
	AgeYears  int
	AgeMonths int
	AgeDays   int
	Zodiac    string
}

func calculateAge(birthDate time.Time) (int, int, int) {
	currentDate := time.Now()

	years := currentDate.Year() - birthDate.Year()
	birthDate = birthDate.AddDate(years, 0, 0)

	if currentDate.Before(birthDate) {
		years--
		birthDate = birthDate.AddDate(-1, 0, 0)
	}

	months := int(currentDate.Sub(birthDate).Hours() / 24 / 30)
	birthDate = birthDate.AddDate(0, months, 0)

	days := int(currentDate.Sub(birthDate).Hours() / 24)

	return years, months, days
}

func getZodiacName(birthDate time.Time, db *sql.DB) (string, error) {
	query := `
        SELECT StartDate, EndDate, ZodiacName
        FROM TZodiak;
    `

	rows, err := db.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var zodiacName string
	for rows.Next() {
		var startDate, endDate time.Time
		var currentZodiacName string
		err := rows.Scan(&startDate, &endDate, &currentZodiacName)
		if err != nil {
			return "", err
		}

		// Compare birthDate within the range of startDate and endDate
		if (birthDate.Month() == startDate.Month() && birthDate.Day() >= startDate.Day()) ||
			(birthDate.Month() == endDate.Month() && birthDate.Day() <= endDate.Day()) {
			zodiacName = currentZodiacName
			break
		}
	}

	if zodiacName == "" {
		return "Unknown", nil
	}

	return zodiacName, nil
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		name := r.FormValue("name")
		birthDateStr := r.FormValue("birth_date")
		birthDate, _ := time.Parse("2006-01-02", birthDateStr)

		years, months, days := calculateAge(birthDate)

		//zodiac := "Unknown" // Implement fetching Zodiac from database here
		zodiac, err := getZodiacName(birthDate, db)
		if err != nil {
			http.Error(w, "Error retrieving zodiac name", http.StatusInternalServerError)
			return
		}

		data := PageData{
			Name:      name,
			AgeYears:  years,
			AgeMonths: months,
			AgeDays:   days,
			Zodiac:    zodiac,
		}

		tmpl := template.Must(template.ParseFiles("template.html"))
		tmpl.Execute(w, data)
		return
	}

	tmpl := template.Must(template.ParseFiles("form.html"))
	tmpl.Execute(w, nil)
}

func main() {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	var err error
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}
