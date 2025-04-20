package commands

import (
	"fmt"
	"time"
	"webserver/src/database"
	"webserver/src/server/sessions"
)

func Plan(session *sessions.Session, args []string) {
	var rank string
	switch session.User.Rank {
	case 0:
		rank = "banned"
		break
	case 1:
		rank = "user"
		break
	case 2:
		rank = "Admin"
		break
	default:
		rank = "Invalid..."
		break
	}
	t := time.Unix(int64(session.User.Expire), 0)
	date := t.Format(time.UnixDate)
	fmt.Fprintf(session.Conn, "[+]==============[+]\r\n")
	fmt.Fprintf(session.Conn, "Username: %s\r\n", session.User.Username)
	fmt.Fprintf(session.Conn, "Api Key: %s\r\n", session.User.ApiKey)
	fmt.Fprintf(session.Conn, "Max Time: %d\r\n", session.User.MaxTime)
	fmt.Fprintf(session.Conn, "Concurrents: %d\r\n", session.User.Conns)
	fmt.Fprintf(session.Conn, "Rank: %s\r\n", rank)
	fmt.Fprintf(session.Conn, "Expire: %s\r\n", date)
	fmt.Fprintf(session.Conn, "Running: %d\r\n", database.GetConcurrents(session.User.Username))
	fmt.Fprintf(session.Conn, "[+]==============[+]\r\n")

}
