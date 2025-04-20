package commands

import (
	"fmt"
	"strings"
	"webserver/src/database"
	"webserver/src/server/sessions"
)

func ShowMethods(session *sessions.Session, args []string) {
	// fetch methods from SSH
	fmt.Fprintf(session.Conn, "[+]########################[+]\r\n")
	methods, err := database.GetMethods()
	if err != nil {
		fmt.Fprintf(session.Conn, "Something Went Wrong Fetching methods!\r\n")
		return
	}
	if methods == nil {
		fmt.Fprintf(session.Conn, "There Are No Methods Available!\r\n")
	}
	for _, v := range methods {
		if v.Active == 1 {
			methodarray := strings.Split(v.Methods, ",")
			for _, m := range methodarray {
				fmt.Fprintf(session.Conn, "            %s\r\n", m)
			}
		}
	}
	fmt.Fprintf(session.Conn, "[+]########################[+]\r\n")
	// fetch methods from DB
}
