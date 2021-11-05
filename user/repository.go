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

func (repo *repo) getUser(id string) (User, error) {
	var u User

	err := repo.conn.QueryRow(
		context.Background(),
		"SELECT id, name, email, phone FROM users WHERE id=$1", id,
	).Scan(
		&u.Id, &u.Name, &u.Email, &u.Phone,
	)
	if err != nil {
		log.Println(err.Error())
	}

	return u, err
}
