package commands

import (
	"fmt"
	"webserver/src/database"
	"webserver/src/server/sessions"
)

func ClearLogs(sesh *sessions.Session, args []string) {
	fmt.Fprintf(sesh.Conn, "Are You Sure You Want To Delete Your Logs(y/n): ")
	check, err := sesh.Term.ReadLine("Are You Sure You Want To Delete Your Logs(y/n): ")
	if err != nil {
		return
	}
	if check == "y" || check == "Y" {
		if err := database.ClearLogs(sesh.User.Username); err != nil {
			fmt.Fprintf(sesh.Conn, "%s\r\n", err)
			return
		} else {
			fmt.Fprintf(sesh.Conn, "Cleared Logs Successfully...\r\n")
			return
		}
	} else {
		return
	}
}
