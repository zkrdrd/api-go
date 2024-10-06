package postgredb

import (
	"api-go/pkg/models"
	"log"
)

// Получение пользователя из БД по id
func (db *DB) GetUser(id string) (*models.Users, error) {
	user := &models.Users{}
	if err := db.conn.QueryRow(`
	SELECT first_name, last_name, middle_name
	FROM customers WHERE id = $1;`, id).Scan(
		&user.FirstName,
		&user.LastName,
		&user.MiddleName); err != nil {
		log.Print(err)
		return nil, err
	}
	return user, nil
}

// Запись пользователя в БД
func (db *DB) SaveUser(user *models.Users) error {
	if _, err := db.conn.Exec(`
	INSERT INTO customers (first_name, last_name, middle_name) 
	VALUES (
	$1, --FirstName
	$2, --LastName
	$3); --MiddleName `,
		user.FirstName,
		user.LastName,
		user.MiddleName); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
