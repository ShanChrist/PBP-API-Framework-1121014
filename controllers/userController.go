package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Fatal(err)
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.UserType)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Fatal(err)
		}
		users = append(users, user)
	}

	sendUserResponse(c, 200, "Success", users)
}

func InsertUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	err := c.Request.ParseForm()
	if err != nil {
		sendResponse(c, http.StatusBadRequest, "Bad Request")
		return
	}

	username := c.PostForm("username")
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	email := c.PostForm("email")
	password := c.PostForm("password")
	usertype := 1

	_, err = db.Exec("INSERT INTO users (username, firstname, lastname, email, password, usertype) values (?,?,?,?,?,?)",
		username,
		firstname,
		lastname,
		email,
		password,
		usertype,
	)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Fatal(err)
		sendResponse(c, 400, "Insert Unsuccesfull!!")
		return
	}
	sendResponse(c, 200, "Success Inserted!!")
}

func UpdateUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	idUser := c.Param("id")

	err := c.Request.ParseForm()
	if err != nil {
		sendResponse(c, http.StatusBadRequest, "Bad Request")
		return
	}

	username := c.PostForm("username")
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	email := c.PostForm("email")

	query := "UPDATE users SET username = ?, firstname = ?, lastname = ?, email = ? WHERE id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		sendResponse(c, 400, "Erorr")
		log.Println(err)
		return
	}

	result, err := stmt.Exec(username, firstname, lastname, email, idUser)
	if err != nil {
		sendResponse(c, 400, "Erorr")

		log.Println(err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sendResponse(c, 400, "Erorr")
		log.Println(err)
		return
	}

	if rowsAffected == 0 {
		sendResponse(c, 400, "The id may not exist in the table.")
		return
	}
	sendResponse(c, 200, "Success Updated")
}

func DeleteUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	idUser := c.Param("id")

	query := "delete from users where id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		sendResponse(c, 400, "Erorr")
		log.Println(err)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(idUser)
	if err != nil {
		sendResponse(c, 400, "Erorr")
		log.Println(err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sendResponse(c, 400, "Erorr")
		log.Println(err)
		return
	}

	if rowsAffected == 0 {
		sendResponse(c, 400, "The id may not exist in the table.")
		return
	}
	sendResponse(c, 200, "Success Deleted")
}

func GetSpecificUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	idUser := c.Param("id")

	query := "SELECT * FROM users WHERE id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		sendResponse(c, http.StatusBadRequest, "Error preparing query")
		log.Println(err)
		return
	}
	defer stmt.Close()

	var user User

	row := stmt.QueryRow(idUser)
	err = row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.UserType)
	if err == sql.ErrNoRows {
		sendResponse(c, http.StatusNotFound, "User not found")
		return
	} else if err != nil {
		sendResponse(c, http.StatusInternalServerError, "Error querying database")
		log.Println(err)
		return
	}

	sendUserResponse(c, http.StatusOK, "Success", []User{user})
}
