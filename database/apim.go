package database

import (
	"database/sql"
	"errors"
)

// APIS
func AddNewAPI(link string, tag string, active int) error {
	rows, err := Sql.Query("INSERT INTO `apis` (link, tag, active) VALUES(?, ?, ?)", link, tag, active)
	if err != nil {
		return err
	}
	if rows != nil {
		rows.Close()
	}
	return err
}

func RemoveAPI(tag string) error {
	rows, err := Sql.Query("DELETE FROM `apis` WHERE `tag` = ?", tag)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Could Not Find API Name")
		}
		return err
	}
	rows.Close()
	return err
}

// SSH SERVERS
func AddSSH(host string, port int, username string, password string, tag string) error {
	rows, err := Sql.Query("INSERT INTO `ssh_servers` (host, port, username, password, tag) VALUES(?, ?, ?, ?, ?)", host, port, username, password, tag)
	if err != nil {
		return err
	}
	if rows != nil {
		rows.Close()
	}
	return err
}

func RemoveSSH(tag string) error {
	rows, err := Sql.Query("DELETE FROM `ssh_servers` WHERE `tag` = ?", tag)
	if err != nil {
		return err
	}
	rows.Close()
	return err
}

// MIRAIS
func AddMirai(host string, port int, username string, password string, tag string) error {
	rows, err := Sql.Query("INSERT INTO `mirais` (host, port, username, password, tag) VALUES(?, ?, ?, ?, ?)", host, port, username, password, tag)
	if err != nil {
		return err
	}
	if rows != nil {
		rows.Close()
	}
	return err
}

func RemoveMirai(tag string) error {
	rows, err := Sql.Query("DELETE FROM `mirais` WHERE `tag` = ?", tag)
	if err != nil {
		return err
	}
	rows.Close()
	return err
}

func AddQBOT(host string, port int, username string, password string, tag string) error {
	rows, err := Sql.Query("INSERT INTO `qbots` (host, port, username, password, tag) VALUES(?, ?, ?, ?, ?)", host, port, username, password, tag)
	if err != nil {
		return err
	}
	if rows != nil {
		rows.Close()
	}
	return err
}

func RemoveQbot(tag string) error {
	rows, err := Sql.Query("DELETE FROM `qbots` WHERE `tag` = ?", tag)
	if err != nil {
		return err
	}
	rows.Close()
	return err
}
