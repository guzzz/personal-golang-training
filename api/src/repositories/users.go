package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Users representa um repositório de usuários
type Users struct {
	db *sql.DB
}

// NewUsersRepository cria repositório de usuários
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {

	statement, erro := repository.db.Prepare(
		"INSERT into USERS (name, nick, email, password) VALUES (?, ?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastInsertId, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertId), nil
}

func (repository Users) List(param string) ([]models.User, error) {
	nameOrNick := fmt.Sprintf("%%%s%%", param)

	rows, erro := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users WHERE name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var tempUser models.User
		if erro = rows.Scan(
			&tempUser.ID,
			&tempUser.Name,
			&tempUser.Nick,
			&tempUser.Email,
			&tempUser.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, tempUser)

	}

	return users, nil
}

func (repository Users) Retrieve(userId uint64) (models.User, error) {

	var tempUser models.User
	linha, erro := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users WHERE id = ?",
		userId,
	)
	if erro != nil {
		return tempUser, erro
	}

	defer linha.Close()

	for linha.Next() {
		if erro = linha.Scan(
			&tempUser.ID,
			&tempUser.Name,
			&tempUser.Nick,
			&tempUser.Email,
			&tempUser.CreatedAt,
		); erro != nil {
			return tempUser, erro
		}
	}

	return tempUser, nil
}

func (repository Users) Update(userId uint64, data models.User) error {

	statement, erro := repository.db.Prepare(
		"UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(data.Name, data.Email, data.Nick, userId)

	if erro != nil {
		return erro
	}
	return nil
}

func (repository Users) Delete(userId uint64) (int64, error) {

	statement, erro := repository.db.Prepare(
		"DELETE FROM users WHERE id = ?;",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(userId)
	if erro != nil {
		return 0, erro
	}

	rows_deleted, erro := result.RowsAffected()
	if erro != nil {
		return 0, erro
	}

	return rows_deleted, nil
}

func (repository Users) RetrieveUserForLogin(email string) (models.User, error) {

	var tempUser models.User
	linha, erro := repository.db.Query(
		"SELECT id, email, password FROM users WHERE email = ?",
		email,
	)
	if erro != nil {
		return tempUser, erro
	}

	defer linha.Close()

	for linha.Next() {
		if erro = linha.Scan(
			&tempUser.ID,
			&tempUser.Email,
			&tempUser.Password,
		); erro != nil {
			return tempUser, erro
		}
	}

	return tempUser, nil
}

func (repository Users) RetrieveUserForUpdatePassword(userId uint64) (models.User, error) {

	var tempUser models.User
	linha, erro := repository.db.Query(
		"SELECT id, email, password FROM users WHERE id = ?",
		userId,
	)
	if erro != nil {
		return tempUser, erro
	}

	defer linha.Close()

	for linha.Next() {
		if erro = linha.Scan(
			&tempUser.ID,
			&tempUser.Email,
			&tempUser.Password,
		); erro != nil {
			return tempUser, erro
		}
	}

	return tempUser, nil
}

func (repository Users) FollowUser(userId, followerId uint64) error {
	statement, erro := repository.db.Prepare(
		"INSERT ignore into followers (user_id, follower_id) VALUES (?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(userId, followerId)
	if erro != nil {
		return erro
	}
	return nil
}

func (repository Users) UnfollowUser(userId, followerId uint64) (bool, error) {
	var success bool = false

	statement, erro := repository.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? and follower_id = ?;",
	)
	if erro != nil {
		return success, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(userId, followerId)
	if erro != nil {
		return success, erro
	}

	rows_afected, erro := result.RowsAffected()
	if erro != nil {
		return success, nil
	}

	if rows_afected > 0 {
		success = true
	}

	return success, nil
}

func (repository Users) ListFollowers(userId uint64) ([]models.User, error) {

	rows, erro := repository.db.Query(
		`
			SELECT u.id, u.name, u.nick, u.email, u.createdAt 
			FROM users u
			INNER JOIN followers f
			ON u.id = f.follower_id
			WHERE f.user_id = ?
		`,
		userId,
	)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var tempUser models.User
		if erro = rows.Scan(
			&tempUser.ID,
			&tempUser.Name,
			&tempUser.Nick,
			&tempUser.Email,
			&tempUser.CreatedAt,
		); erro != nil {
			return nil, erro
		}
		users = append(users, tempUser)
	}

	return users, nil
}

func (repository Users) ListFollowing(userId uint64) ([]models.User, error) {

	rows, erro := repository.db.Query(
		`
			SELECT u.id, u.name, u.nick, u.email, u.createdAt 
			FROM users u
			INNER JOIN followers f
			ON u.id = f.user_id
			WHERE f.follower_id = ?
		`,
		userId,
	)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var tempUser models.User
		if erro = rows.Scan(
			&tempUser.ID,
			&tempUser.Name,
			&tempUser.Nick,
			&tempUser.Email,
			&tempUser.CreatedAt,
		); erro != nil {
			return nil, erro
		}
		users = append(users, tempUser)
	}

	return users, nil
}

func (repository Users) UpdatePassword(userId uint64, password string) error {

	statement, erro := repository.db.Prepare(
		"UPDATE users SET password = ? WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(password, userId)

	if erro != nil {
		return erro
	}
	return nil
}
