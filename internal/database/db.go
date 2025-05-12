package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Robert076/quack-stack.git/internal/duck"
	_ "github.com/lib/pq"
)

func GetDucksFromDatabase() ([]duck.Duck, error) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error opening db connection: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging db: %v", err)
	}

	query := `
		SELECT
			"Id",
			"Name",
			"Species",
			"Age",
			"Weight",
			"Color",
			"IsMigratory",
			"CanFly",
			"QuackVolume",
			"Mood",
			"FavoriteFood",
			"IsLeader",
			"CreatedAt"
		FROM "Duck";
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying db: %v", err)
	}
	defer rows.Close()

	var ducks []duck.Duck

	for rows.Next() {
		var fetchedDuck duck.Duck
		err := rows.Scan(
			&fetchedDuck.Id,
			&fetchedDuck.Name,
			&fetchedDuck.Species,
			&fetchedDuck.Age,
			&fetchedDuck.Weight,
			&fetchedDuck.Color,
			&fetchedDuck.IsMigratory,
			&fetchedDuck.CanFly,
			&fetchedDuck.QuackVolume,
			&fetchedDuck.Mood,
			&fetchedDuck.FavoriteFood,
			&fetchedDuck.IsLeader,
			&fetchedDuck.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		ducks = append(ducks, fetchedDuck)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return ducks, nil
}
