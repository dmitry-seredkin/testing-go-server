package user

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type repo struct {
	conn *pgx.Conn
}

func InitializeRepository(conn *pgx.Conn) *repo {
	return &repo{conn}
}

func (repo *repo) createUser(cu CreateUser) (User, error) {
	var u User
	err := repo.conn.QueryRow(
		context.Background(),
		"INSERT INTO users (name, email, phone) VALUES ($1, $2, $3) RETURNING id, name, email, phone",
		&cu.Name, &cu.Email, &cu.Phone,
	).Scan(
		&u.Id, &u.Name, &u.Email, &u.Phone,
	)
	if err != nil {
		log.Println(err.Error())
	}

	return u, err
}

func (repo *repo) deleteUser(id string) error {
	var returnedId string
	err := repo.conn.QueryRow(
		context.Background(),
		"DELETE FROM users WHERE id=$1 RETURNING id",
		id,
	).Scan(&returnedId)

	return err
}

func (repo *repo) getUser(id string) (User, error) {
	var u User

	err := repo.conn.QueryRow(
		context.Background(),
		"SELECT id, name, email, phone FROM users WHERE id=$1",
		id,
	).Scan(
		&u.Id, &u.Name, &u.Email, &u.Phone,
	)
	if err != nil {
		log.Println(err.Error())
	}

	return u, err
}

func (repo *repo) getUsers() ([]UserItem, error) {
	users := []UserItem{}

	rows, err := repo.conn.Query(
		context.Background(),
		"SELECT id, name FROM users",
	)
	if err != nil {
		log.Println(err.Error())
	} else {
		var u UserItem

		for rows.Next() {
			if err := rows.Scan(&u.Id, &u.Name); err == nil {
				users = append(users, u)
			}
		}
	}

	return users, err
}
