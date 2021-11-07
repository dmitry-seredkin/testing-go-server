package user

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type repo struct {
	conn        *pgx.Conn
	userService UserService
}

func InitializeRepository(conn *pgx.Conn) *repo {
	return &repo{conn, UserService{}}
}

func (repo *repo) createUser(user CreateUser) (User, error) {
	hash, _ := repo.userService.hashPassword(user.Password)

	var u User
	err := repo.conn.QueryRow(
		context.Background(),
		`
			INSERT INTO users (login, name, password, email, phone)
			VALUES ($1, $2, $3, $4, $5) 
			RETURNING id, login, name, email, phone
		`,
		&user.Login, &user.Name, hash, &user.Email, &user.Phone,
	).Scan(
		&u.Id, &u.Login, &u.Name, &u.Email, &u.Phone,
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

func (repo *repo) loginUser(user LoginUser) (bool, error) {
	var hash string
	err := repo.conn.QueryRow(
		context.Background(),
		"SELECT password FROM users WHERE login=$1",
		user.Login,
	).Scan(&hash)

	return repo.userService.checkPassword(user.Password, hash), err
}

func (repo *repo) updateUser(id string, user UpdateUser) (User, error) {
	var u User
	err := repo.conn.QueryRow(
		context.Background(),
		`
			UPDATE users SET
				name = COALESCE($2, name),
				email = COALESCE($3, email),
				phone = COALESCE($4, phone)
			WHERE id=$1
			RETURNING id, login, name, email, phone
		`,
		id, &user.Name, &user.Email, &user.Phone,
	).Scan(
		&u.Id, &u.Login, &u.Name, &u.Email, &u.Phone,
	)
	if err != nil {
		log.Println(err.Error())
	}

	return u, err
}
