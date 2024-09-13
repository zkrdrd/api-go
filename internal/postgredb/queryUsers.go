package postgredb

import (
	"api-go/internal/buisness"
	"log"
)

func (db *DB) GetUser(user *buisness.Users) (*buisness.Users, error) {
	if err := db.Conn.QueryRow(`
	SELECT id first_name, last_name, middle_name
	FROM transaction`).Scan(
		&user.FirstName,
		&user.LastName,
		&user.MiddleName); err != nil {
		log.Print(err)
		return nil, err
	}
	return user, nil
}

func (db *DB) SaveUser(transf *buisness.Users) error {
	if _, err := db.Conn.Exec(`
	INSERT INTO transaction (first_name, last_name, middle_name) 
	VALUES (
	?, -- FirstName
    ?, -- LastName
    ?  -- MiddleName)`,
		transf.FirstName,
		transf.LastName,
		transf.MiddleName); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
