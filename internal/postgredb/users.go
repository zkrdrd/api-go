package postgredb

import (
	"github.com/zkrdrd/api-go/pkg/models"
)

// Получение пользователя из БД по id.
func (db *DB) GetUser(id string) (*models.Users, error) {
	user := &models.Users{}
	if err := db.useConn().QueryRow(`
	SELECT first_name, last_name, middle_name
	FROM customers WHERE id = $1;`, id).Scan(
		&user.FirstName,
		&user.LastName,
		&user.MiddleName); err != nil {
		return nil, err
	}
	return user, nil
}

// Запись пользователя в БД.
func (db *DB) SaveUser(user *models.Users) error {
	if _, err := db.useConn().Exec(`
	INSERT INTO customers (first_name, last_name, middle_name) 
	VALUES (
	$1, --FirstName
	$2, --LastName
	$3); --MiddleName `,
		user.FirstName,
		user.LastName,
		user.MiddleName); err != nil {
		return err
	}
	return nil
}
