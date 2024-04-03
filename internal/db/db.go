package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDB(dbName string) (*sql.DB, error) {
	var dbs *sql.DB
	dbs, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	return dbs, err
}
func CloseDB(dbs *sql.DB) {
	dbs.Close()
}

func CreateMigrationsTable(dbs *sql.DB) {

	_, err := dbs.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`)
	if err != nil {
		log.Fatal("error in creating users table : ", err)
	}
	_, err = dbs.Exec(`CREATE TABLE user_info (
		user_id INTEGER PRIMARY KEY,
		user_rank TEXT DEFAULT 'beginner' CHECK (user_rank IN ('beginner', 'intermediate', 'advanced', 'expert')),
		user_points INTEGER DEFAULT 1000,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	  );`)
	if err != nil {
		log.Fatal("error in creating user_info table : ", err)
	}
}

func CreateUser(dbs *sql.DB, users Users) (user Users, userInfo UserInfo, err error) {
	var u Users
	var ui UserInfo

	res, err := dbs.Exec(`INSERT INTO users (username, email, password) VALUES (?, ?, ?)`, users.Username, users.Email, users.Password)
	if err != nil {
		log.Fatal("error in creating user : ", err)
	}
	// get last insert id
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal("error in getting last inserted userId")
	}
	res, err = dbs.Exec(`INSERT INTO user_info (user_id) VALUES (?)`, lastID)
	if err != nil {
		log.Fatal("error in creating user info : ", lastID, err)
	}
	uiLastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal("error in getting last inserted userId from usersinfo")
	}
	err = dbs.QueryRow("SELECT * FROM users WHERE id = $1", lastID).Scan(&u.Id, &u.Username, &u.Email, &u.Password)
	if err != nil {
		log.Fatal("error in getting user : ", err)
	}
	err = dbs.QueryRow("SELECT * FROM user_info WHERE user_id = $1", uiLastId).Scan(&ui.User_id, &ui.Username, &ui.UserRank, &ui.UserRating)
	if err != nil {
		log.Fatal("error in getting user info : ", err)
	}
	return u, ui, err
}
