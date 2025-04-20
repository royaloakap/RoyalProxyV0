package commands

import (
	"fmt"
	"log"
	"strings"
	"time"
	"webserver/src/database"
	"webserver/src/server/sessions"

	"github.com/alexeyco/simpletable"
)

func AllAttacks(sesh *sessions.Session, args []string) {
	if sesh.User.Rank != 2 {
		fmt.Fprintf(sesh.Conn, "Insufficient Permissions\r\n")
		return
	} else {
		table := simpletable.New()
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "#"},
				{Align: simpletable.AlignCenter, Text: "Username"},
				{Align: simpletable.AlignCenter, Text: "Host"},
				{Align: simpletable.AlignCenter, Text: "Port"},
				{Align: simpletable.AlignCenter, Text: "Duration"},
				{Align: simpletable.AlignCenter, Text: "Method"},
				{Align: simpletable.AlignCenter, Text: "Time Left"},
			},
		}
		attacks, err := database.GetAllRunning()
		if err != nil {
			log.Printf("(err/db) %s", err)
			return
		}

		if attacks == nil {
			fmt.Fprintf(sesh.Conn, "No Attacks Are Running...\r\n")
			return
		}

		for i, attack := range attacks {
			tmleft := (attack.Duration + int(attack.TimeSent)) - int(time.Now().Unix())
			r := []*simpletable.Cell{
				{Align: simpletable.AlignRight, Text: fmt.Sprint(i + 1)},
				{Text: fmt.Sprint(attack.Username)},
				{Text: fmt.Sprint(attack.Host)},
				{Text: fmt.Sprint(attack.Port)},
				{Text: fmt.Sprint(attack.Duration)},
				{Text: fmt.Sprint(attack.Method)},
				{Text: fmt.Sprintf("%ds", tmleft)},
			}
			table.Body.Cells = append(table.Body.Cells, r)
		}
		table.SetStyle(simpletable.StyleCompactClassic)
		fmt.Fprintf(sesh.Conn, strings.Replace("    "+table.String(), "\n", "\r\n    ", -1)+"\r\n")
	}
}
