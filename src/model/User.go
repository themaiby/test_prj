// contains methods for access DB tables
package model

import log "github.com/kataras/golog"
import (
	"../db"
	"strconv"
)

// must be defined in every controller
var table = "users"

// struct for users table
type user struct {
	ID         int     `db:"id" json:"id"`
	Username   *string `db:"username" json:"username"`
	Name       *string `db:"name" json:"name"`
	First_name *string `db:"first_name" json:"first_name"`
	Last_name  *string `db:"last_name" json:"last_name"`
	Role_ID    *int    `db:"role_id" json:"role_id"`
}

// Model var for export
var User = user{}

// get all rows
func (this *user) All() *[]user {
	// prepare new var for users
	users := []user{}

	// check connection
	if !db.IsConnected() {
		log.Info("DB wasn't connected. Reconnecting...")
		db.Reconnect()
	}

	// query
	err := db.DB.Select(&users, "SELECT id, username, name, first_name, last_name, role_id from "+table)

	// check errors
	if err != nil {
		log.Warn(err)
	}
	return &users
}

// get single row by id
func (this *user) Find(id int) *user {
	// prepare new var for users
	user := user{}

	// query
	result := db.DB.QueryRowx("SELECT id, username, name, first_name, last_name, role_id FROM "+table+" where id = ?", id)

	// scan query result into struct var
	result.StructScan(&user)
	return &user
}

// get first row
func (this *user) First() *user {
	// prepare new var for users
	user := user{}

	result := db.DB.QueryRowx("SELECT id, username, name, first_name, last_name, role_id from " + table + " LIMIT")
	// scan query result into struct var
	result.StructScan(&user)
	return &user
}

// get TOP rows TODO: need to find usage
func (this *user) Limit(limit int) *[]user {
	users := []user{}
	// generating condition
	err := db.DB.Select(&users, "SELECT TOP "+strconv.Itoa(limit)+" id, username, name, first_name, last_name, role_id from "+table)
	if err != nil {
		log.Error("[MSSQL] : ", err)
	}
	return &users
}
