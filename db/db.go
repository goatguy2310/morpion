package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DBConn struct {
	DB *sql.DB
}

func NewDBConnection(path string) DBConn {
	newDB, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}

	// create tables if not exist

	// user_handles table
	_, err = newDB.Exec(
		`CREATE TABLE IF NOT EXISTS user_handles (
		user_id 	TEXT 	NOT NULL 	PRIMARY KEY,
		handle 		TEXT 	NOT NULL
		);`,
	)
	if err != nil {
		panic(err)
	}

	// challenge table
	_, err = newDB.Exec(
		`CREATE TABLE IF NOT EXISTS challenges (
		user_id 	TEXT 	NOT NULL 	PRIMARY KEY,
		opponent_id TEXT  NOT NULL,
		criteria 	BLOB
		);`,
	)
	if err != nil {
		panic(err)
	}

	// duel table
	_, err = newDB.Exec(
		`CREATE TABLE IF NOT EXISTS duels (
		user1_id 	TEXT 	NOT NULL 	PRIMARY KEY,
		user2_id 	TEXT  	NOT NULL,
		handle1 	TEXT 	NOT NULL,
		handle2 	TEXT 	NOT NULL,
		game_state 	BLOB
		);`,
	)
	if err != nil {
		panic(err)
	}

	return DBConn{DB: newDB}
}

func (db DBConn) UpdateUserHandle(userID string, handle string) error {
	_, err := db.DB.Exec(
		`INSERT OR REPLACE INTO
			user_handles (user_id, handle)
		VAlUES 
			(?, ?);`,
		userID,
		handle,
	)
	return err
}

func (db DBConn) GetUserHandle(userID string) (string, error) {
	row := db.DB.QueryRow(
		`SELECT handle
		FROM user_handles
		WHERE user_id = ?;`,
		userID,
	)

	handle := ""

	err := row.Scan(&handle)
	return handle, err
}

func (db DBConn) UpdateChallenge(userID string, opponentID string, criteria []byte) error {
	_, err := db.DB.Exec(
		`INSERT OR REPLACE INTO
			challenges (user_id, opponent_id, criteria)
		VALUES
			(?, ?, ?);`,
		userID,
		opponentID,
		criteria,
	)
	return err
}

func (db DBConn) GetChallengeByUserID(userID string) (string, []byte, error) {
	row := db.DB.QueryRow(
		`SELECT opponent_id, criteria
		FROM challenges
		WHERE user_id = ?;`,
		userID,
	)

	var opponentID string
	var criteria []byte

	err := row.Scan(&opponentID, &criteria)
	return opponentID, criteria, err
}

func (db DBConn) GetChallengeByOpponentID(opponentID string) (string, []byte, error) {
	row := db.DB.QueryRow(
		`SELECT user_id, criteria
		FROM challenges
		WHERE opponent_id = ?;`,
		opponentID,
	)

	var userID string
	var criteria []byte

	err := row.Scan(&userID, &criteria)
	return userID, criteria, err
}

func (db DBConn) DeleteChallenge(userID string) error {
	_, err := db.DB.Exec(
		`DELETE FROM challenges
		WHERE user_id = ?;`,
		userID,
	)
	return err
}

func (db DBConn) UpdateDuel(user1ID string, user2ID string, handle1 string, handle2 string, gameState []byte) error {
	_, err := db.DB.Exec(
		`INSERT OR REPLACE INTO
			duels (user1_id, user2_id, handle1, handle2, game_state)
		VALUES
			(?, ?, ?, ?, ?);`,
		user1ID,
		user2ID,
		handle1,
		handle2,
		gameState,
	)
	return err
}

func (db DBConn) GetDuelByUserID(userID string) (string, string, string, string, []byte, error) {
	row := db.DB.QueryRow(
		`SELECT user2_id, handle2, game_state
		FROM duels
		WHERE user1_id = ?;`,
		userID,
	)

	var user2ID string
	var handle2 string
	var gameState []byte

	err := row.Scan(&user2ID, &handle2, &gameState)
	return user2ID, handle2, user2ID, handle2, gameState, err
}

func (db DBConn) GetDuelByOpponentID(opponentID string) (string, string, string, string, []byte, error) {
	row := db.DB.QueryRow(
		`SELECT user1_id, handle1, game_state
		FROM duels
		WHERE user2_id = ?;`,
		opponentID,
	)

	var user1ID string
	var handle1 string
	var gameState []byte

	err := row.Scan(&user1ID, &handle1, &gameState)
	return user1ID, handle1, user1ID, handle1, gameState, err
}

func (db DBConn) DeleteDuel(userID string) error {
	_, err := db.DB.Exec(
		`DELETE FROM duels
		WHERE user1_id = ?;`,
		userID,
	)
	return err
}

func (db DBConn) Close() {
	db.DB.Close()
}
