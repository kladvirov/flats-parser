package flats

import (
	"flats-parser/db"
	"fmt"
	"log"
	"strings"
)

type Flat struct {
	ID        int64
	RemoteID  int
	Type      int8
	CreatedAt string
	UpdatedAt string
}

func Get[ID int | string](t int8, IDs []ID) []Flat {
	var flats []Flat

	placeholders := make([]string, len(IDs))
	args := make([]interface{}, len(IDs)+1)
	args[0] = t

	for i, id := range IDs {
		placeholders[i] = "?"
		args[i+1] = id
	}

	query := fmt.Sprintf("SELECT * FROM flats WHERE type = ? AND remote_id IN (%s)", strings.Join(placeholders, ","))
	rows, err := db.Db.Query(query, args...)

	if err != nil {
		log.Fatal(err)
		return []Flat{}
	}
	defer rows.Close()

	for rows.Next() {
		var flat Flat
		if err := rows.Scan(&flat.ID, &flat.RemoteID, &flat.Type, &flat.CreatedAt, &flat.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		flats = append(flats, flat)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return []Flat{}
	}

	return flats
}

func Insert(flats []Flat) error {
	if len(flats) == 0 {
		return nil
	}

	query := "INSERT INTO flats (type, remote_id) VALUES "
	values := make([]interface{}, 0, len(flats)*2)

	for i, flat := range flats {
		if i > 0 {
			query += ","
		}
		query += "(?, ?)"
		values = append(values, flat.Type, flat.RemoteID)
	}

	_, err := db.Db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}
