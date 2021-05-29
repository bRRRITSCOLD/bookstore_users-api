package users_domain

import (
	users_mysql_db "bookstore_users-api/datasources/users/mysql"
	"bookstore_users-api/logger"
	dates_utils "bookstore_users-api/utils/dates"
	errors_utils "bookstore_users-api/utils/errors"
	"fmt"
)

const (
	USERS_MYSQL_DB_INSERT_USER_QUERY                       = "INSERT INTO users(firstName, lastName, email, dateCreated, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	USERS_MYSQL_DB_PUT_USER_BY_ID_QUERY                    = "UPDATE users SET firstName=?, lastName=?, email=?, status=?, password=? WHERE id=?;"
	USERS_MYSQL_DB_SELECT_USER_BY_ID_QUERY                 = "SELECT * from users WHERE id=?;"
	USERS_MYSQL_DB_SELECT_USERS_BY_STATUS_QUERY            = "SELECT * from users WHERE status=?;"
	USERS_MYSQL_DB_SELECT_USER_BY_EMAIL_AND_PASSWORD_QUERY = "SELECT * from users WHERE email=? AND password=? AND status=?;"
	USERS_MYSQL_DB_DELETE_USER_BY_ID_QUERY                 = "DELETE FROM users WHERE id=?"
	USERS_MYSQL_DB_EMAIL_UNIQUE                            = "email_UNIQUE"
	USERS_MYSQL_DB_NO_ROWS                                 = "sql: no rows in result set"
)

func (user *User) GetByUserID() *errors_utils.APIError {
	stmt, prepareErr := users_mysql_db.Client.Preparex(USERS_MYSQL_DB_SELECT_USER_BY_ID_QUERY)
	if prepareErr != nil {
		logger.Error("error when trying to prepare USERS_MYSQL_DB_SELECT_USER_BY_ID_QUERY statement", prepareErr)
		return errors_utils.NewInternalServerAPIError("internal server error")
	}

	defer stmt.Close()

	var foundUser User

	queryRowResult := stmt.QueryRowx(user.UserID)
	if scanStructErr := queryRowResult.StructScan(&foundUser); scanStructErr != nil {
		logger.Error("error when scanning struct for USERS_MYSQL_DB_SELECT_USER_BY_ID_QUERY result", scanStructErr)
		return errors_utils.NewInternalServerAPIError("internal server error")
	}

	user.UserID = foundUser.UserID
	user.Email = foundUser.Email
	user.FirstName = foundUser.FirstName
	user.LastName = foundUser.LastName
	user.DateCreated = foundUser.DateCreated
	user.Status = foundUser.Status

	return nil
}

func (user *User) GetUserByEmailAndPassword() *errors_utils.APIError {
	stmt, prepareErr := users_mysql_db.Client.Preparex(USERS_MYSQL_DB_SELECT_USER_BY_EMAIL_AND_PASSWORD_QUERY)
	if prepareErr != nil {
		logger.Error("error when trying to prepare USERS_MYSQL_DB_SELECT_USER_BY_EMAIL_AND_PASSWORD_QUERY statement", prepareErr)
		return errors_utils.NewInternalServerAPIError("internal server error")
	}

	defer stmt.Close()

	var foundUser User

	queryRowResult := stmt.QueryRowx(user.Email, user.Password, USER_ACTIVE_STATUS)
	if scanStructErr := queryRowResult.StructScan(&foundUser); scanStructErr != nil {
		logger.Error("error when scanning struct for USERS_MYSQL_DB_SELECT_USER_BY_ID_QUERY result", scanStructErr)
		return errors_utils.ParseMySQLError(scanStructErr)
	}

	user.UserID = foundUser.UserID
	user.Email = foundUser.Email
	user.FirstName = foundUser.FirstName
	user.LastName = foundUser.LastName
	user.DateCreated = foundUser.DateCreated
	user.Status = foundUser.Status

	return nil
}

func (user *User) Save() *errors_utils.APIError {
	stmt, prepareErr := users_mysql_db.Client.Preparex(USERS_MYSQL_DB_INSERT_USER_QUERY)
	if prepareErr != nil {
		logger.Error("error when trying to prepare USERS_MYSQL_DB_INSERT_USER_QUERY statement", prepareErr)
		return errors_utils.NewInternalServerAPIError("internal server error")
	}

	defer stmt.Close()

	insertResult, insertErr := stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
		user.Status,
		user.Password,
	)
	if insertErr != nil {
		logger.Error("error when inserting row for USERS_MYSQL_DB_INSERT_USER_QUERY", insertErr)
		return errors_utils.NewInternalServerAPIError("internal server error")
	}

	userId, lastInsertedIDErr := insertResult.LastInsertId()
	if lastInsertedIDErr != nil {
		logger.Error("error getting last inserted id for USERS_MYSQL_DB_INSERT_USER_QUERY", lastInsertedIDErr)
		return errors_utils.NewInternalServerAPIError("internal server error")
	}

	user.UserID = userId

	return nil
}

func (user *User) PutByUserID() *errors_utils.APIError {
	stmt, prepareErr := users_mysql_db.Client.Preparex(USERS_MYSQL_DB_PUT_USER_BY_ID_QUERY)
	if prepareErr != nil {
		logger.Error("error when trying to prepare USERS_MYSQL_DB_PUT_USER_BY_ID_QUERY statement", prepareErr)
		return errors_utils.NewInternalServerAPIError("internal server error")
	}

	defer stmt.Close()

	user.DateCreated = dates_utils.GetNow()

	_, putErr := stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.UserID,
		user.Status,
		user.Password,
	)
	if putErr != nil {
		logger.Error("error when putting row for USERS_MYSQL_DB_PUT_USER_BY_ID_QUERY", putErr)
		return errors_utils.NewInternalServerAPIError("internal server error")
	}

	return nil
}

func (user *User) DeleteByUserID() *errors_utils.APIError {
	stmt, prepareErr := users_mysql_db.Client.Preparex(USERS_MYSQL_DB_DELETE_USER_BY_ID_QUERY)
	if prepareErr != nil {
		logger.Error("error when trying to prepare USERS_MYSQL_DB_DELETE_USER_BY_ID_QUERY statement", prepareErr)
		return errors_utils.NewInternalServerAPIError("internal server error")
	}

	defer stmt.Close()

	if _, deleteErr := stmt.Exec(user.UserID); deleteErr != nil {
		logger.Error("error when deleting row for USERS_MYSQL_DB_DELETE_USER_BY_ID_QUERY", deleteErr)
		return errors_utils.NewInternalServerAPIError("internal server error")
	}

	return nil
}

func (user *User) GetUsersByStatus(status string) (Users, *errors_utils.APIError) {
	stmt, prepareErr := users_mysql_db.Client.Preparex(USERS_MYSQL_DB_SELECT_USERS_BY_STATUS_QUERY)
	if prepareErr != nil {
		logger.Error("error when trying to prepare USERS_MYSQL_DB_SELECT_USERS_BY_STATUS_QUERY statement", prepareErr)
		return nil, errors_utils.NewInternalServerAPIError("internal server error")
	}

	defer stmt.Close()

	queryxRows, queryxErr := stmt.Queryx(status)
	if queryxErr != nil {
		logger.Error("error when running query for USERS_MYSQL_DB_SELECT_USERS_BY_STATUS_QUERY", queryxErr)
		return nil, errors_utils.NewInternalServerAPIError("internal server error")
	}

	defer queryxRows.Close()

	foundUsers := make([]User, 0)

	for queryxRows.Next() {
		var foundUser User

		if scanStructErr := queryxRows.StructScan(&foundUser); scanStructErr != nil {
			logger.Error("error when scanning stuct for USERS_MYSQL_DB_SELECT_USERS_BY_STATUS_QUERY", queryxErr)
			return nil, errors_utils.NewInternalServerAPIError("internal server error")
		}

		foundUsers = append(foundUsers, foundUser)
	}

	if len(foundUsers) == 0 {
		return nil, errors_utils.NewNotFoundAPIError(fmt.Sprintf("no users found matching status %s", status))
	}

	return foundUsers, nil
}
