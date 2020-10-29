package handlers
import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/testApi/config"
	"github.com/testApi/models"
	"net/http"
)

// GETHandler ...
func GETHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-Id")
	db := config.OpenConnection()
	var person models.Person
	row := db.QueryRow("SELECT user_id, first_name, last_name FROM person where user_id = $1", userID)
	err := row.Scan(
		&person.FirstName,
		&person.LastName,
		&person.UserID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rows, err := db.Query(`
			SELECT address from person_addresses where user_id = $1`, userID,
	)
	for rows.Next() {
		var address models.Address

		err = rows.Scan(
			&address.Address,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		return
		}

		person.Addresses = append(person.Addresses, address)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	

	personBytes, _ := json.MarshalIndent(person, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(personBytes)
	defer rows.Close()
	defer db.Close()
}

// POSTHandler ...
func POSTHandler(w http.ResponseWriter, r *http.Request) {
	db := config.OpenConnection()

	var p models.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID, err := uuid.NewRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO person (user_id, first_name, last_name) VALUES ($1, $2, $3)`
	_, err = tx.Exec(sqlStatement, userID, p.FirstName, p.LastName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	for _, address := range p.Addresses {
		insertNew := `INSERT INTO person_addresses (user_id, address) VALUES ($1, $2)`
		_, err = tx.Exec(
			insertNew,
			userID,
			address.Address,
		)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusBadRequest)
		return
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}