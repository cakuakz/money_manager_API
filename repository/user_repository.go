package repository

import (
	"context"
	"database/sql"
	"money-manager/entity"
	"money-manager/structs"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUser(db *sql.DB) (usr structs.GetAllUserPaginationResponse, err error) {
	query := "SELECT * FROM users"

	rows, err := db.Query(query)
	if err != nil {
		return usr, err
	}

	defer rows.Close()

	for rows.Next() {
		var user entity.Users

		err := rows.Scan(&user.UserID, &user.Username, &user.Fullname, &user.CreatedAt, &user.CreatedAt)
		if err != nil {
			return usr, err
		}

		usr.Users = append(usr.Users, user)
	}

	return usr, nil
}

func RegisterUser(db *sql.DB, data structs.UserCreateRequest) error {
	query := "INSERT INTO users (username, fullname, password) VALUES ($1, $2,$3)"

	errs := db.QueryRow(query, data.Username, data.Fullname, data.Password)

	return errs.Err()
}

func UpdateUser(db *sql.DB, data structs.UserUpdateRequest) error {
	query := "UPDATE users SET username = $1, fullname = $2 WHERE id = $3"

	errs := db.QueryRow(query, data.Username, data.Fullname, data.UserID)

	return errs.Err()
}

func DeletePerson(db *sql.DB, data structs.UserDeletionRequest) (err error) {
    sql := "DELETE FROM person WHERE id = $1"

    errs := db.QueryRow(sql, data.UserID)
    return errs.Err()
}

func VerifyUser(ctx context.Context, db *sql.DB, user structs.UserLoginRequest) (usr entity.Users, err error) {
	query := "SELECT user_id, username, password, fullname FROM users WHERE username = $1"

	row := db.QueryRowContext(ctx, query, user.Username)
	if err := row.Scan(&usr.UserID, &usr.Username, &usr.Password, &usr.Fullname); err != nil {
		if err == sql.ErrNoRows {
			return entity.Users{}, errors.New("user not found")
		}
		return entity.Users{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(user.Password)); err != nil {
		return entity.Users{}, errors.New("invalid password")
	}

	usr.Password = ""
	return usr, nil
}