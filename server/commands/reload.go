package commands

import (
	"fmt"
	"log"
	"webserver/src/config"
	"webserver/src/server/sessions"
)

func Reload(sesh *sessions.Session, args []string) {
	if sesh.User.Rank != 2 {
		fmt.Fprintf(sesh.Conn, "Insuffcient Permissions\r\n")
		return
	} else {
		if err := config.LoadConfig("config.json"); err != nil {
			fmt.Fprintf(sesh.Conn, "Something Went Wrong Reloading Configuration File...\r\n")
			log.Printf("(config_reload_err) %s", err)
			return
		} else {
			fmt.Fprintf(sesh.Conn, "Successfully Reloaded File\r\n")
		}
	}
}
